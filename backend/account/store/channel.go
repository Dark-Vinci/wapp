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
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

type Channel struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source channel.go -destination ./mock/channel_mock.go -package mock ChannelDatabase
type ChannelDatabase interface {
	Create(ctx context.Context, channel account.Channel) (*account.Channel, error)
	DeleteByID(ctx context.Context, channelID uuid.UUID, deletedAt time.Time) error
	GetChannelByID(ctx context.Context, channelID uuid.UUID) (*account.Channel, error)
	DeleteUserChannels(ctx context.Context, userID uuid.UUID, deletedAt time.Time, tx *gorm.DB) error
}

func NewChannel(conn *Store) *ChannelDatabase {
	log := conn.Log.With().
		Str(constants.FunctionNameHelper, "NewChannel").
		Str(constants.PackageStrHelper, packageName).
		Logger()

	channel := &Channel{
		logger: &log,
		db:     conn.Connection,
	}

	channelDB := ChannelDatabase(channel)

	return &channelDB
}

func (c *Channel) Create(ctx context.Context, channel account.Channel) (*account.Channel, error) {
	log := c.logger.With().
		Str(constants.MethodStrHelper, "channel.Create").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Got request to create channel %+v", channel)

	if res := c.db.WithContext(ctx).Model(&account.Channel{}).Create(&channel); res.Error != nil {
		log.Err(res.Error).Msg("Failed to create channel")

		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return nil, sdkerror.ErrDuplicateKey
		}

		return nil, sdkerror.ErrRecordCreation
	}

	return &channel, nil
}

func (c *Channel) GetChannelByID(ctx context.Context, channelID uuid.UUID) (*account.Channel, error) {
	log := c.logger.With().
		Str(constants.MethodStrHelper, "channel.GetChannelByID").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	var channel account.Channel

	if res := c.db.WithContext(ctx).Model(&account.Channel{}).Where("id = ?", channelID).First(&channel); res.Error != nil {
		log.Err(res.Error).Msg("Failed to get channel")

		return nil, sdkerror.ErrRecordNotFound
	}

	return &channel, nil
}

func (c *Channel) DeleteUserChannels(ctx context.Context, userID uuid.UUID, deletedAt time.Time, tx *gorm.DB) error {
	log := c.logger.With().
		Str(constants.MethodStrHelper, "channel.DeleteUserChannels").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Got request to delete user channels %v", userID)

	if tx != nil {
		log.Info().Msg("deleting user channels in a transaction")

		if err := tx.WithContext(ctx).Model(&account.Channel{}).Where("user_id = ?", userID).UpdateColumns(&account.Channel{DeletedAt: &deletedAt}).Error; err != nil {
			log.Err(err).Msg("Failed to delete user channels")
			return err
		}

		return nil
	}

	if err := c.db.WithContext(ctx).Model(&account.Channel{}).Where("user_id = ?", userID).UpdateColumns(&account.Channel{DeletedAt: &deletedAt}).Error; err != nil {
		log.Err(err).Msg("Failed to delete user channels")
		return err
	}

	return nil
}

func (c *Channel) DeleteByID(ctx context.Context, channelID uuid.UUID, deletedAt time.Time) error {
	log := c.logger.With().
		Str(constants.MethodStrHelper, "channel.DeleteByID").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Got request to delete channel with ID %v", channelID)

	if res := c.db.WithContext(ctx).Model(&account.Channel{}).Where("id = ?", channelID).UpdateColumns(account.Channel{DeletedAt: &deletedAt}); res.Error != nil {
		log.Err(res.Error).Msg("Failed to delete channel")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToDeleteRecord
	}

	return nil
}
