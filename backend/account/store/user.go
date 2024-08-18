package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/account/connection"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

type UserDatabase interface {
}

type User struct {
	log        *zerolog.Logger
	connection *gorm.DB
}

func NewUser(db *connection.DBConn) *UserDatabase {
	u := User{
		connection: db.Connection,
		log:        db.Log,
	}

	uu := UserDatabase(u)

	return &uu
}

func (u *User) CreateUser(ctx context.Context, user models.UserModel) (*models.UserModel, error) {
	log := u.log.With().
		Str(constants.MethodStrHelper, "store.CreateUser").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.PackageStrHelper, packageName).
		Logger()

	log.Info().Msg("Got request to create user")

	res := u.connection.Create(&user)

	if res.Error != nil {
		log.Err(res.Error).Msg("failed to create user")
		return nil, res.Error
	}

	return &user, nil
}

func (u *User) GetUserByID(ctx context.Context, id uuid.UUID) (*models.UserModel, error) {
	log := u.log.With().
		Str(constants.PackageStrHelper, packageName).
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "store.GetUserByID").
		Logger()

	log.Info().Msg("Got request to get user by id")

	return &models.UserModel{}, nil
}
