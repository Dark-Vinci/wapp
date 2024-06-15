package store

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/linkedout/backend/account/connection"
)

type UserReader interface{}
type UserWriter interface{}

type UserDatabase interface {
	UserReader
	UserWriter
}

type User struct {
	log            *zerolog.Logger
	master         *gorm.DB
	slaves         []*gorm.DB
	count          int64
	slaveInstances int
}

func NewUser(db *connection.DBConn) *UserDatabase {
	u := User{
		master:         db.Master,
		slaves:         db.Slaves,
		log:            db.Log,
		count:          int64(0),
		slaveInstances: len(db.Slaves),
	}

	uu := UserDatabase(u)

	return &uu
}

type U struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	DateOfBirth time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAT   *time.Time
}

func (u *User) CreateUser(ctx context.Context, user U) (*U, error) {
	log := u.log.With().
		Str("KEY", "store.CreateUser").
		Logger()

	log.Info().Msg("Got request to create user")

	res := u.master.Create(&user)

	if res.Error != nil {
		log.Err(res.Error).Msg("failed to create user")
		return nil, res.Error
	}

	return &user, nil
}

func (u *User) getSlaveConnection() *gorm.DB {
	if len(u.slaves) == 0 {
		return u.master
	}

	index := u.count % int64(u.slaveInstances)

	u.count++

	return u.slaves[index]
}

func (u *User) GetUserByID(ctx context.Context, id uuid.UUID) (*U, error) {
	log := u.log.With().
		Str("KEY", "store.GetUserByID").
		Logger()

	log.Info().Msg("Got request to get user by id")

	conn := u.getSlaveConnection()

	res := U{}

	r := conn.Model(U{ID: id}).First(&res)

	if r.Error != nil {
		log.Err(r.Error).Msg("Unable to fetch user by Id")
		return nil, r.Error
	}

	return &res, nil
}
