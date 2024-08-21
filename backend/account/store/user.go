package store

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/account/connection"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models"
	"github.com/dark-vinci/wapp/backend/sdk/sdkerror"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

//go:generate mockgen -source user.go -destination ./mock/user_mock.go -package mock UserDatabase
type UserDatabase interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	Delete(ctx context.Context, id uuid.UUID, deletedAt time.Time, tx *gorm.DB) error
}

type User struct {
	log        *zerolog.Logger
	connection *gorm.DB
}

func NewUser(db *connection.DBConn) *UserDatabase {
	l := db.Log.With().
		Str(constants.FunctionNameHelper, "NewUser").
		Str(constants.PackageStrHelper, packageName).
		Logger()

	user := &User{
		connection: db.Connection,
		log:        &l,
	}

	operations := UserDatabase(user)

	return &operations
}

func (u *User) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	log := u.log.With().
		Str(constants.MethodStrHelper, "store.CreateUser").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Got a request to create user")

	if res := u.connection.WithContext(ctx).Model(&models.User{}).Create(&user); res.Error != nil {
		log.Err(res.Error).Msg("error creating user")

		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return nil, sdkerror.ErrDuplicateKey
		}

		return nil, sdkerror.ErrRecordCreation
	}

	return &user, nil
}

func (u *User) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	log := u.log.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "store.GetUserByID").
		Logger()

	log.Info().Msg("Got a request to get user by id")

	var user models.User

	if res := u.connection.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Find(&user); res.Error != nil {
		log.Err(res.Error).Msg("unable to get user by id")
		return nil, sdkerror.ErrRecordNotFound
	}

	return &user, nil
}

func (u *User) Delete(ctx context.Context, id uuid.UUID, deletedAt time.Time, tx *gorm.DB) error {
	log := u.log.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "user.DeleteUser").
		Logger()

	log.Info().Msg("Got a request to delete user by id")

	if tx != nil {
		if res := tx.WithContext(ctx).Where("id = ?", id).UpdateColumns(models.User{DeletedAt: &deletedAt}); res.Error != nil {
			log.Err(res.Error).Msg("unable to delete user")

			if errors.Is(res.Error, gorm.ErrRecordNotFound) {
				return sdkerror.ErrRecordNotFound
			}

			return sdkerror.ErrFailedToDeleteRecord
		}
	}

	if res := u.connection.WithContext(ctx).Where("id = ?", id).UpdateColumns(models.User{DeletedAt: &deletedAt}); res.Error != nil {
		log.Err(res.Error).Msg("unable to delete user")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToDeleteRecord
	}

	return nil
}
