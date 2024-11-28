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

type SpecifiedUserMultiUserChatMedia struct {
	logger zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source specified_user_multi_user_media.go -destination ./mock/specified_user_multi_user_media_mock.go package mock SpecifiedUserMultiUserChatMediaDatabase
type SpecifiedUserMultiUserChatMediaDatabase interface {
	GetByID(ctx context.Context, entityID uuid.UUID) (*media.SpecifiedUserMultiUserChatMedia, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewSpecifiedUserMultiUserChatMedia(logger zerolog.Logger) SpecifiedUserMultiUserChatMediaDatabase {
	return &SpecifiedUserMultiUserChatMedia{}
}

func (p *SpecifiedUserMultiUserChatMedia) GetByID(ctx context.Context, entityID uuid.UUID) (*media.SpecifiedUserMultiUserChatMedia, error) {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "specifiedUserMultiUserChatMedia.GetByID").Logger()

	var specifiedUserMultiUserChatMedia media.SpecifiedUserMultiUserChatMedia

	res := p.db.WithContext(ctx).Model(&media.SpecifiedUserMultiUserChatMedia{}).Where("id = ? and seen = false", entityID).First(&specifiedUserMultiUserChatMedia)
	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to get SpecifiedUserMultiUserChatMedia")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, sdkerror.ErrRecordNotFound
		}

		return nil, sdkerror.ErrRecordCreation // todo: update accordingly
	}

	return &specifiedUserMultiUserChatMedia, nil
}

func (p *SpecifiedUserMultiUserChatMedia) Delete(ctx context.Context, id uuid.UUID) error {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "SpecifiedUserMultiUserChatMedia.Delete").Logger()

	res := p.db.WithContext(ctx).Model(&media.SpecifiedUserMultiUserChatMedia{}).UpdateColumns(media.SpecifiedUserMultiUserChatMedia{ID: id})

	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to delete SpecifiedUserMultiUserChatMedia")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrSomethingWentWrong
	}

	return nil
}
