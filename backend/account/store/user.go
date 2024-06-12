package store

import (
	"context"
	"sync/atomic"
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
	log    *zerolog.Logger
	master *gorm.DB
	slaves []*gorm.DB
}

func NewUser(db *connection.DBConn) *UserDatabase {
	u := User{
		master: db.Master,
		slaves: db.Slaves,
		log:    db.Log,
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

func (u *User) GetUserByID(ctx context.Context, id uuid.UUID) (*U, error) {
	log := u.log.With().
		Str("KEY", "store.GetUserByID").
		Logger()

	log.Info().Msg("Got request to get user by id")

	resChan := make(chan U)
	errChan := make(chan error)
	var errorCount atomic.Uint64

	for _, v := range u.slaves {
		go func(resChan chan<- U, v *gorm.DB) {
			res := U{}
			r := v.Model(U{ID: id}).First(&res)

			if r.Error != nil {
				errorCount.Add(1)
				errChan <- r.Error
				log.Err(r.Error).Msg("Unable to fetch user by Id")
				return
			}

			resChan <- res
		}(resChan, v)
	}

	

	select {
	case user := <-resChan:
		return &user, nil
	case e := <- errChan:
		return nil, e
	}
}
