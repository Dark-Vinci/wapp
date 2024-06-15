package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/linkedout/backend/account/connection"
	"github.com/dark-vinci/linkedout/backend/sdk/constants"
	"github.com/dark-vinci/linkedout/backend/sdk/models"
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

func (u *User) CreateUser(ctx context.Context, user models.UserModel) (*models.UserModel, error) {
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

func (u *User) GetUserByID(ctx context.Context, id uuid.UUID) (*models.UserModel, error) {
	log := u.log.With().
		Str(constants.PackageStrHelper, packageName).
		Str(constants.FunctionNameHelper, "store.GetUserByID").
		Logger()

	log.Info().Msg("Got request to get user by id")

	conn := u.getSlaveConnection()

	res := models.UserModel{}

	r := conn.Model(models.UserModel{ID: id}).First(&res)

	if r.Error != nil {
		log.Err(r.Error).Msg("Unable to fetch user by Id")
		return nil, r.Error
	}

	return &res, nil
}
