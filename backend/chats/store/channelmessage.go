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

type ChannelMessage struct {
	db     *gorm.DB
	logger *zerolog.Logger
}

//go:generate mockgen -source channelmessage.go -destination ./mock/channelmessage_mock.go -package mock ChannelMessageDatabase
type ChannelMessageDatabase interface {
	Delete(ctx context.Context, id uuid.UUID, deletedAt time.Time) error
	Fetch(ctx context.Context, channelID uuid.UUID) ([]chats.ChannelMessage, error)
	Create(ctx context.Context, chat chats.ChannelMessage) (*chats.ChannelMessage, error)
	GetMessageByID(ctx context.Context, id, userID uuid.UUID) (*chats.ChannelMessage, error)
}

func NewChannelMessageDatabase(conn *connections.DBConn) *ChannelMessageDatabase {
	logger := conn.Logger.With().
		Str(constants.FunctionNameHelper, "NewChannelMessageDatabase").
		Str(constants.PackageStrHelper, packageName).
		Logger()

	channelMessage := &ChannelMessage{
		db:     conn.Connection,
		logger: conn.Logger,
	}

	channelMessageOperations := ChannelMessageDatabase(channelMessage)

	logger.Info().Msg("Successfully created new channel chat database")

	return &channelMessageOperations
}

func (cc *ChannelMessage) Create(ctx context.Context, chat chats.ChannelMessage) (*chats.ChannelMessage, error) {
	log := cc.logger.With().
		Str(constants.MethodStrHelper, "channelMessage.Create").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Creating new channel chat")

	if err := cc.db.WithContext(ctx).Model(&chats.ChannelMessage{}).Create(&chat).Error; err != nil {
		log.Err(err).Msg("Failed to create new channel chat")

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, sdkerror.ErrDuplicateKey
		}

		return nil, sdkerror.ErrDuplicateKey
	}

	return &chat, nil
}

func (cc *ChannelMessage) Delete(ctx context.Context, id uuid.UUID, deletedAt time.Time) error {
	log := cc.logger.With().
		Str(constants.MethodStrHelper, "channelMessage.Delete").
		Str(constants.RequestID, id.String()).
		Logger()

	log.Info().Msg("Deleting new channel chat")

	if err := cc.db.WithContext(ctx).Model(&chats.ChannelMessage{}).Where("id = ?", id).UpdateColumns(&chats.ChannelMessage{DeletedAt: &deletedAt}).Error; err != nil {
		log.Err(err).Msg("Failed to delete new channel chat")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrRecordNotFound
	}

	return nil
}

func (cc *ChannelMessage) GetMessageByID(ctx context.Context, id, userID uuid.UUID) (*chats.ChannelMessage, error) {
	log := cc.logger.With().
		Str(constants.FunctionNameHelper, "channelMessage.GetMessageByID").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Got a request get channel chat by id")

	var result chats.ChannelMessage

	if err := cc.db.WithContext(ctx).Model(&chats.ChannelMessage{}).Where("id = ? and user_id = ?", id, userID).First(&result).Error; err != nil {
		log.Err(err).Msg("Failed to get channel chat by id")
		return nil, sdkerror.ErrRecordNotFound
	}

	return &result, nil
}

func (cc *ChannelMessage) Fetch(ctx context.Context, channelID uuid.UUID) ([]chats.ChannelMessage, error) {
	log := cc.logger.With().
		Str(constants.MethodStrHelper, "channelMessage.Fetch").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	var result []chats.ChannelMessage

	if err := cc.db.WithContext(ctx).Model(&chats.ChannelMessage{}).Where("channel_id = ?", channelID).Find(&result).Error; err != nil {
		log.Err(err).Msg("Failed to get channel chat")
		return nil, sdkerror.ErrRecordNotFound
	}

	return result, nil
}
