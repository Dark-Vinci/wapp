package store

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models/account"
	"github.com/dark-vinci/wapp/backend/sdk/sdkerror"
)

type UserPassword struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source userpassword.go -destination ./mock/userpassword_mock.go -package mock UserPasswordDatabase
type UserPasswordDatabase interface {
	Create(ctx context.Context, userPassword account.UserPasswords) (*account.UserPasswords, error)
	Update(ctx context.Context, userPassword account.UserPasswords) error
	Delete(ctx context.Context, deletedAt time.Time, passwordID, userID uuid.UUID) error
	GetUserPassword(ctx context.Context, userID uuid.UUID) ([]account.UserPasswords, error)
}

func NewUserPassword(conn *Store) *UserPasswordDatabase {
	logger := conn.Log.With().
		Str(constants.FunctionNameHelper, "NewUserPassword").
		Str(constants.PackageStrHelper, packageName).Logger()

	userPassword := &UserPassword{
		logger: &logger,
		db:     conn.Connection,
	}

	logger.Info().Msg("successfully created new user password database")

	userPasswordOperations := UserPasswordDatabase(userPassword)

	return &userPasswordOperations
}

func (up *UserPassword) Create(ctx context.Context, userPassword account.UserPasswords) (*account.UserPasswords, error) {
	log := up.logger.With().
		Str(constants.MethodStrHelper, "userPassword.Create").
		Str(constants.PackageStrHelper, packageName).Logger()

	log.Info().Msg("creating new user password database")

	if err := up.db.WithContext(ctx).Model(&account.UserPasswords{}).Create(&userPassword).Error; err != nil {
		log.Err(err).Msg("failed to create user password database")

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, sdkerror.ErrDuplicateKey
		}

		return nil, sdkerror.ErrRecordCreation
	}

	return &userPassword, nil
}

func (up *UserPassword) Update(ctx context.Context, userPassword account.UserPasswords) error {
	log := up.logger.With().
		Str(constants.MethodStrHelper, "userPassword.Update").
		Str(constants.PackageStrHelper, packageName).Logger()

	log.Info().Msg("updating user password database")

	if err := up.db.WithContext(ctx).Model(&account.UserPasswords{}).Save(&userPassword).Error; err != nil {
		log.Err(err).Msg("failed to update user password database")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToUpdateRecord
	}

	return nil
}

func (up *UserPassword) GetUserPassword(ctx context.Context, userID uuid.UUID) ([]account.UserPasswords, error) {
	log := up.logger.With().
		Str(constants.MethodStrHelper, "userPassword.GetUserPassword").
		Str(constants.PackageStrHelper, packageName).Logger()

	log.Info().Msg("getting user password database")

	var result []account.UserPasswords

	if err := up.db.WithContext(ctx).Model(&account.UserPasswords{}).Where("user_id = ?", userID).Find(&result).Error; err != nil {
		log.Err(err).Msg("failed to get user password database")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sdkerror.ErrRecordNotFound
		}

		return nil, sdkerror.ErrRecordNotFound // todo: update
	}

	return result, nil
}

func (up *UserPassword) Delete(ctx context.Context, deletedAt time.Time, passwordID, userID uuid.UUID) error {
	log := up.logger.With().
		Str(constants.MethodStrHelper, "userPassword.Delete").
		Str(constants.PackageStrHelper, packageName).Logger()

	log.Info().Msg("deleting user password database")

	if err := up.db.WithContext(ctx).Model(&account.UserPasswords{}).Where("id = ? and user_id = ?", passwordID, userID).UpdateColumns(&account.UserPasswords{DeletedAt: &deletedAt}).Error; err != nil {
		log.Err(err).Msg("failed to delete user password database")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToDeleteRecord
	}

	return nil
}
