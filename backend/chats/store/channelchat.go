package store

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/chats/connections"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models/chats"
	"github.com/dark-vinci/wapp/backend/sdk/sdkerror"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

type ChannelChat struct {
	db     *gorm.DB
	logger *zerolog.Logger
}

type ChannelChatDatabase interface {
	Delete(ctx context.Context, id uuid.UUID, deletedAt time.Time) error
	Fetch(ctx context.Context, channelID uuid.UUID) ([]chats.ChannelChat, error)
	Create(ctx context.Context, chat chats.ChannelChat) (*chats.ChannelChat, error)
	GetChatByID(ctx context.Context, id, userID uuid.UUID) (*chats.ChannelChat, error)
}

func NewChannelChatDatabase(conn *connections.DBConn) *ChannelChatDatabase {
	logger := conn.Logger.With().
		Str(constants.FunctionNameHelper, "NewChannelChatDatabase").
		Str(constants.PackageStrHelper, packageName).
		Logger()

	channelChat := &ChannelChat{
		db:     conn.Connection,
		logger: conn.Logger,
	}

	channelChatOperations := ChannelChatDatabase(channelChat)

	logger.Info().Msg("Successfully created new channel chat database")

	return &channelChatOperations
}

func (cc *ChannelChat) Create(ctx context.Context, chat chats.ChannelChat) (*chats.ChannelChat, error) {
	log := cc.logger.With().
		Str(constants.MethodStrHelper, "channelChat.Create").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Creating new channel chat")

	if err := cc.db.WithContext(ctx).Model(&chats.ChannelChat{}).Create(&chat).Error; err != nil {
		log.Err(err).Msg("Failed to create new channel chat")

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, sdkerror.ErrDuplicateKey
		}

		return nil, sdkerror.ErrDuplicateKey
	}

	return &chat, nil
}

func (cc *ChannelChat) Delete(ctx context.Context, id uuid.UUID, deletedAt time.Time) error {
	log := cc.logger.With().
		Str(constants.MethodStrHelper, "channelChat.Delete").
		Str(constants.RequestID, id.String()).
		Logger()

	log.Info().Msg("Deleting new channel chat")

	if err := cc.db.WithContext(ctx).Model(&chats.ChannelChat{}).Where("id = ?", id).UpdateColumns(&chats.ChannelChat{DeletedAt: &deletedAt}).Error; err != nil {
		log.Err(err).Msg("Failed to delete new channel chat")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrRecordNotFound
	}

	return nil
}

func (cc *ChannelChat) GetChatByID(ctx context.Context, id, userID uuid.UUID) (*chats.ChannelChat, error) {
	log := cc.logger.With().
		Str(constants.FunctionNameHelper, "channelChat.GetChatByID").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Got a request get channel chat by id")

	var result chats.ChannelChat

	if err := cc.db.WithContext(ctx).Model(&chats.ChannelChat{}).Where("id = ? and user_id = ?", id, userID).First(&result).Error; err != nil {
		log.Err(err).Msg("Failed to get channel chat by id")
		return nil, sdkerror.ErrRecordNotFound
	}

	return &result, nil
}

func (cc *ChannelChat) Fetch(ctx context.Context, channelID uuid.UUID) ([]chats.ChannelChat, error) {
	log := cc.logger.With().
		Str(constants.MethodStrHelper, "channelChat.Fetch").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	var result []chats.ChannelChat

	if err := cc.db.WithContext(ctx).Model(&chats.ChannelChat{}).Where("channel_id = ?", channelID).Find(&result).Error; err != nil {
		log.Err(err).Msg("Failed to get channel chat")
		return nil, sdkerror.ErrRecordNotFound
	}

	return result, nil
}
