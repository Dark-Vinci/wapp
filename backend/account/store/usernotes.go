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
	"github.com/dark-vinci/wapp/backend/sdk/models/account"
	"github.com/dark-vinci/wapp/backend/sdk/sdkerror"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

type UserNotes struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source usernotes.go -destination ./mock/usernotes_mock.go -package mock UserNoteDatabase
type UserNoteDatabase interface {
	Create(ctx context.Context, userNote account.UserNotes) (*account.UserNotes, error)
	Update(ctx context.Context, userNote account.UserNotes) error
	Delete(ctx context.Context, deletedAt time.Time, userID, noteID uuid.UUID) error
}

func NewUserNote(conn *connection.DBConn) *UserNoteDatabase {
	logger := conn.Log.With().
		Str(constants.FunctionNameHelper, "NewUserNote").
		Str(constants.PackageStrHelper, packageName).
		Logger()

	userNote := &UserNotes{
		logger: &logger,
		db:     conn.Connection,
	}

	userNoteOperations := UserNoteDatabase(userNote)

	return &userNoteOperations
}

func (u *UserNotes) Create(ctx context.Context, userNote account.UserNotes) (*account.UserNotes, error) {
	log := u.logger.With().
		Str(constants.MethodStrHelper, "userNote.Create").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Got a reques to create user not with payload %v", userNote)

	if err := u.db.WithContext(ctx).Model(&account.UserNotes{}).Create(&userNote).Error; err != nil {
		log.Err(err).Msg("Failed to create userNote")

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, sdkerror.ErrDuplicateKey
		}

		return nil, err
	}

	return &userNote, nil
}

func (u *UserNotes) Update(ctx context.Context, userNote account.UserNotes) error {
	log := u.logger.With().
		Str(constants.MethodStrHelper, "userNote.Update").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Got a request to update user not with payload %v", userNote)

	if err := u.db.WithContext(ctx).Model(&account.UserNotes{}).Where("id = ?", userNote.ID).UpdateColumns(&userNote).Error; err != nil {
		log.Info().Msg("Failed to update userNote")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToUpdateRecord
	}

	return nil
}

func (u *UserNotes) Delete(ctx context.Context, deletedAt time.Time, userID, noteID uuid.UUID) error {
	log := u.logger.With().
		Str(constants.MethodStrHelper, "userNote.DeleteUser").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Got a request to delete the user %s, note of id %s", userID, noteID)

	if err := u.db.WithContext(ctx).Where("user_id = ? and id = ?", userID, noteID).UpdateColumns(&account.UserNotes{DeletedAt: &deletedAt}).Error; err != nil {
		log.Err(err).Msg("Failed to update user")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToDeleteRecord
	}

	return nil
}
