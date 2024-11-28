package store

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models/media"
	"github.com/dark-vinci/wapp/backend/sdk/sdkerror"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

type ChatMedia struct {
	logger zerolog.Logger
	db     *gorm.DB
}

type ChatMediaDatabase interface {
	Create(ctx context.Context, chatMedia media.ChatMedia) (*media.ChatMedia, error)
	FetchMessageMedia(ctx context.Context, entityID uuid.UUID) ([]media.ChatMedia, error)
	GetByID(ctx context.Context, entityID uuid.UUID) (*media.ChatMedia, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewChatMedia(logger zerolog.Logger) ChatMediaDatabase {
	return &ChatMedia{}
}

func (p *ChatMedia) GetByID(ctx context.Context, entityID uuid.UUID) (*media.ChatMedia, error) {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "chatMedia.GetByID").Logger()

	var chatMedia media.ChatMedia

	res := p.db.WithContext(ctx).Model(&media.ChatMedia{}).Where("id = ?", entityID).First(&chatMedia)
	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to get chatMedia")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, sdkerror.ErrRecordNotFound
		}

		return nil, sdkerror.ErrRecordCreation // todo: update accordingly
	}

	return &chatMedia, nil
}

func (p *ChatMedia) FetchMessageMedia(ctx context.Context, messageID uuid.UUID) ([]media.ChatMedia, error) {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "chatMedia.FetchMessageMedia").Logger()

	var messageMedia []media.ChatMedia

	res := p.db.WithContext(ctx).Model(&media.ChatMedia{}).Where("message_id = ?", messageID).Find(&messageMedia)
	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to get chatMedia")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, sdkerror.ErrRecordNotFound
		}

		return nil, sdkerror.ErrRecordCreation // todo: update accordingly
	}

	return messageMedia, nil
}

func (p *ChatMedia) Create(ctx context.Context, chatMedia media.ChatMedia) (*media.ChatMedia, error) {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "chatMedia.Create").Logger()

	res := p.db.WithContext(ctx).Model(&media.ChatMedia{}).Create(&chatMedia)

	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to create chatMedia")

		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return nil, sdkerror.ErrDuplicateKey
		}

		return nil, sdkerror.ErrRecordCreation
	}

	return &chatMedia, nil
}

func (p *ChatMedia) Delete(ctx context.Context, id uuid.UUID) error {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "chatMedia.Delete").Logger()

	res := p.db.WithContext(ctx).Model(&media.ChatMedia{}).UpdateColumns(media.ChatMedia{ID: id})

	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to delete chatMedia")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrSomethingWentWrong
	}

	return nil
}
