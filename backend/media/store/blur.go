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

type Blur struct {
	logger zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source blur.go -destination ./mock/blur_mock.go -package mock BlurDatabase
type BlurDatabase interface {
	Create(ctx context.Context, blur media.Blur) (*media.Blur, error)
	GetByID(ctx context.Context, entityID uuid.UUID) (*media.Blur, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewBlur(logger zerolog.Logger, db *gorm.DB) BlurDatabase {
	return &Blur{}
}

func (p *Blur) GetByMediaID(ctx context.Context, mediaID uuid.UUID) (*media.Blur, error) {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "blur.GetByID").Logger()

	var blur media.Blur

	res := p.db.WithContext(ctx).Model(&media.Blur{}).Where("media_id = ?", mediaID).First(&blur)
	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to get blur")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, sdkerror.ErrRecordNotFound
		}

		return nil, sdkerror.ErrRecordCreation // todo: update accordingly
	}

	return &blur, nil
}

func (p *Blur) GetByID(ctx context.Context, id uuid.UUID) (*media.Blur, error) {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "blur.GetByID").Logger()

	var blur media.Blur

	res := p.db.WithContext(ctx).Model(&media.Blur{}).Where("id = ?", id).First(&blur)
	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to get blur")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, sdkerror.ErrRecordNotFound
		}

		return nil, sdkerror.ErrRecordCreation // todo: update accordingly
	}

	return &blur, nil
}

func (p *Blur) Create(ctx context.Context, blur media.Blur) (*media.Blur, error) {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "blur.Create").Logger()

	res := p.db.WithContext(ctx).Model(&media.Blur{}).Create(&blur)

	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to create blur")

		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return nil, sdkerror.ErrDuplicateKey
		}

		return nil, sdkerror.ErrRecordCreation
	}

	return &blur, nil
}

func (p *Blur) Delete(ctx context.Context, id uuid.UUID) error {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "blur.Delete").Logger()

	res := p.db.WithContext(ctx).Model(&media.Blur{}).UpdateColumns(media.Blur{ID: id})

	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to delete blur")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrSomethingWentWrong
	}

	return nil
}
