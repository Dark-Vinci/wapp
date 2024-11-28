package store

// FOR GROUPS AND CHANNELS

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

type MultiUserChatMedia struct {
	logger zerolog.Logger
	db     *gorm.DB
}

type MultiUserChatMediaOperations interface {
	Create(ctx context.Context, multiUserChatMedia media.MultiUserChatMedia) (*media.MultiUserChatMedia, error)
	GetByID(ctx context.Context, entityID uuid.UUID) (*media.MultiUserChatMedia, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewMultiUserChatMedia(logger zerolog.Logger) MultiUserChatMediaOperations {
	return &MultiUserChatMedia{}
}

func (p *MultiUserChatMedia) GetByID(ctx context.Context, entityID uuid.UUID) (*media.MultiUserChatMedia, error) {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "multiUserChatMedia.GetByID").Logger()

	var multiUserChatMedia media.MultiUserChatMedia

	res := p.db.WithContext(ctx).Model(&media.MultiUserChatMedia{}).Where("id = ?", entityID).First(&multiUserChatMedia)
	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to get multiUserChatMedia")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, sdkerror.ErrRecordNotFound
		}

		return nil, sdkerror.ErrRecordCreation // todo: update accordingly
	}

	return &multiUserChatMedia, nil
}

func (p *MultiUserChatMedia) Create(ctx context.Context, multiUserChatMedia media.MultiUserChatMedia) (*media.MultiUserChatMedia, error) {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "multiUserChatMedia.Create").Logger()

	res := p.db.WithContext(ctx).Model(&media.MultiUserChatMedia{}).Create(&multiUserChatMedia)

	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to create multiUserChatMedia")

		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return nil, sdkerror.ErrDuplicateKey
		}

		return nil, sdkerror.ErrRecordCreation
	}

	return &multiUserChatMedia, nil
}

func (p *MultiUserChatMedia) Delete(ctx context.Context, id uuid.UUID) error {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "multiUserChatMedia.Delete").Logger()

	res := p.db.WithContext(ctx).Model(&media.MultiUserChatMedia{}).UpdateColumns(media.MultiUserChatMedia{ID: id})

	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to delete multiUserChatMedia")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrSomethingWentWrong
	}

	return nil
}
