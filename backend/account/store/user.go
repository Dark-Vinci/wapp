package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/account/connection"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models"
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

type UserH[T, M any] struct {
	master *gorm.DB
	slaves []*gorm.DB
}

func NewUserH[T, M any]() *UserH[T, M] {
	return &UserH[T, M]{}
}

const FIRST = "FIRST"

func (u *UserH[T, M]) Modifier(c context.Context, model T, ty string) (*T, error) {
	switch ty {
	case "CREATE":
		db := u.master.WithContext(c).Create(&model)

		if db.Error != nil {
			return nil, db.Error
		}

	case "UPDATE":
		db := u.master.WithContext(c).Updates(&model)

		if db.Error != nil {
			return nil, db.Error
		}
	}

	return &model, nil
}

func (u *UserH[T, M]) Queries(cont context.Context, queryType string) (*T, error) {
	if len(u.slaves) == 0 {
		return nil, errors.New("")
	}
	result := make(chan T)
	errorChan := make(chan error, len(u.slaves))

	c, cancel := context.WithCancel(cont)

	for i, slave := range u.slaves {
		go func(ctx context.Context, index int, result chan<- T, errorChan chan<- error) {
			select {
			case <-ctx.Done():
				return
			default:
			}

			var res T
			var model M

			switch queryType {
			case FIRST:
				if r := slave.Model(model).First(&res); r.Error != nil {
					errorChan <- r.Error
					fmt.Println("ERROR")
					return
				}
			}

			select {
			case <-ctx.Done():
				return
			default:
				result <- res
			}
		}(c, i, result, errorChan)
	}

	var resultReceived T
	var errorCount int

	var currentError error

loop:
	for {
		select {
		case resultReceived = <-result:
			fmt.Println(resultReceived)
			cancel()
			break loop
		case currentError = <-errorChan:
			errorCount++
			if errorCount == len(u.slaves) {
				break loop
			}
		}
	}

	cancel()

	if errorCount == len(u.slaves) {
		return nil, currentError
	}

	return &resultReceived, nil
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

	q := NewUserH[models.UserModel, models.UserModel]()

	res, err := q.Queries(ctx, FIRST)

	if err != nil {
		log.Error().Err(err).Msg("failed to get user by id")
		return nil, err
	}

	return res, nil
}
