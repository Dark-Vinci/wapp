package store

import (
	"context"
	"fmt"
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

func query(db *connection.DBConn) (*struct{}, error) {
	result := make(chan struct{})
	errorChan := make(chan error, len(db.Slaves))

	c, cancel := context.WithCancel(context.Background())

	for i := range db.Slaves {
		go func(ctx context.Context, db *connection.DBConn, index int, result chan<- struct{}, errorChan chan<- error) {
			select {
			case <-ctx.Done():
				return
			default:
			}

			res := models.UserModel{}

			if r := db.Slaves[index].Model(User{}).First(&res); r.Error != nil {
				errorChan <- r.Error
				fmt.Println("ERROR")
				return
			}

			select {
			case <-ctx.Done():
				return
			default:
				result <- struct{}{}
			}
		}(c, db, i, result, errorChan)
	}

	var resultReceived struct{}
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
			if errorCount == len(db.Slaves) {
				break loop
			}
		}
	}

	cancel()

	if errorCount == len(db.Slaves) {
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

	conn := u.getSlaveConnection()

	res := models.UserModel{}

	r := conn.Model(models.UserModel{ID: id}).First(&res)

	if r.Error != nil {
		log.Err(r.Error).Msg("Unable to fetch user by Id")
		return nil, r.Error
	}

	return &res, nil
}
