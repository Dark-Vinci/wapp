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

type PostMedia struct {
	logger zerolog.Logger
	db     *gorm.DB
}

type PostMediaDatabase interface {
	Create(ctx context.Context, postMedia media.PostMedia) (*media.PostMedia, error)
	GetByPostID(ctx context.Context, postID uuid.UUID) (*media.PostMedia, error)
	GetByID(ctx context.Context, postID uuid.UUID) (*media.PostMedia, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewPostMedia(logger zerolog.Logger, db *gorm.DB) PostMediaDatabase {
	return &PostMedia{
		logger,
		db,
	}
}

func (p *PostMedia) GetByID(ctx context.Context, id uuid.UUID) (*media.PostMedia, error) {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "postMedia.GetByID").Logger()

	var postMedia media.PostMedia

	res := p.db.WithContext(ctx).Model(&media.PostMedia{}).Where("id = ?", id).First(&postMedia)
	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to get postMedia")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, sdkerror.ErrRecordNotFound
		}

		return nil, sdkerror.ErrRecordCreation // todo: update accordingly
	}

	return &postMedia, nil
}

func (p *PostMedia) GetByPostID(ctx context.Context, postID uuid.UUID) (*media.PostMedia, error) {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "postMedia.GetByPostID").Logger()

	var postMedia media.PostMedia

	res := p.db.WithContext(ctx).Model(&media.PostMedia{}).Where("entity_id = ?", postID).First(&postMedia)
	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to get postMedia")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, sdkerror.ErrRecordNotFound
		}

		return nil, sdkerror.ErrRecordCreation // todo: update accordingly
	}

	return &postMedia, nil
}

func (p *PostMedia) Create(ctx context.Context, postMedia media.PostMedia) (*media.PostMedia, error) {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "postMedia.Create").Logger()

	res := p.db.WithContext(ctx).Model(&media.PostMedia{}).Create(&postMedia)

	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to create postMedia")

		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return nil, sdkerror.ErrDuplicateKey
		}

		return nil, sdkerror.ErrRecordCreation
	}

	return &postMedia, nil
}

func (p *PostMedia) Delete(ctx context.Context, id uuid.UUID) error {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "postMedia.Delete").Logger()

	res := p.db.WithContext(ctx).Model(&media.PostMedia{}).UpdateColumns(media.PostMedia{ID: id})

	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to delete postMedia")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrSomethingWentWrong
	}

	return nil
}
