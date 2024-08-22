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

type ChannelUser struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source channel_user.go -destination ./mock/channel_user_mock.go -package mock ChannelUserDatabase
type ChannelUserDatabase interface {
	Create(ctx context.Context, channelUser account.ChannelUser) (*account.ChannelUser, error)
	DeleteAllUser(ctx context.Context, channelID uuid.UUID, now time.Time) error
	BlockUser(ctx context.Context, now time.Time, channelID, userID uuid.UUID) error
	UnBlockUser(ctx context.Context, now time.Time, channelID, userID uuid.UUID) error
}

func NewChannelUser(conn *connection.DBConn) *ChannelUserDatabase {
	logger := conn.Log.With().
		Str(constants.FunctionNameHelper, "NewChannelUser").
		Str(constants.PackageStrHelper, packageName).Logger()

	channelUser := &ChannelUser{
		logger: &logger,
		db:     conn.Connection,
	}

	operations := ChannelUserDatabase(channelUser)

	return &operations
}

func (cu *ChannelUser) Create(ctx context.Context, channelUser account.ChannelUser) (*account.ChannelUser, error) {
	log := cu.logger.With().
		Str(constants.MethodStrHelper, "channelUser.Create").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msgf("Got a request to create a new user for a channel with details %v", channelUser)

	if err := cu.db.WithContext(ctx).Model(&account.ChannelUser{}).Create(&channelUser).Error; err != nil {
		log.Err(err).Msg("Failed to create a new user for a channel")

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, sdkerror.ErrDuplicateKey
		}

		return nil, sdkerror.ErrRecordCreation
	}

	return &channelUser, nil
}

func (cu *ChannelUser) DeleteAllUser(ctx context.Context, channelID uuid.UUID, now time.Time) error {
	log := cu.logger.With().
		Str(constants.MethodStrHelper, "channelUser.DeleteAllUser").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msgf("Got a request to delete all the users for a channel with id %v", channelID)

	if err := cu.db.WithContext(ctx).Model(&account.ChannelUser{}).Where("channel_id = ?", channelID).UpdateColumns(&account.ChannelUser{DeletedAt: &now}).Error; err != nil {
		log.Err(err).Msg("Failed to delete all users for a channel")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToDeleteRecord
	}

	return nil
}

func (cu *ChannelUser) BlockUser(ctx context.Context, now time.Time, channelID, userID uuid.UUID) error {
	log := cu.logger.With().
		Str(constants.MethodStrHelper, "channelUser.BlockUser").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msgf("Got a request to block a user for a channel with id %v", channelID)

	if err := cu.db.WithContext(ctx).Model(&account.ChannelUser{}).Where("channel_id = ? and user_id = ?", channelID, userID).UpdateColumns(&account.ChannelUser{Blocked: true, UpdatedAt: now}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToUpdateRecord
	}

	return nil
}

func (cu *ChannelUser) UnBlockUser(ctx context.Context, now time.Time, channelID, userID uuid.UUID) error {
	log := cu.logger.With().
		Str(constants.MethodStrHelper, "channelUser.UnBlockUser").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msgf("Got a request to unblock a user for a channel with id %v", channelID)

	if err := cu.db.WithContext(ctx).Model(&account.ChannelUser{}).Where("channel_id = ? and user_id = ?", channelID, userID).UpdateColumns(&account.ChannelUser{Blocked: false, UpdatedAt: now}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToUpdateRecord
	}

	return nil
}
