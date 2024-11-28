package store

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models/media"
	"github.com/dark-vinci/wapp/backend/sdk/sdkerror"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

type Profile struct {
	logger zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source profile.go -destination ./mock/profile_mock.go -package mock ProfileDatabase
type ProfileDatabase interface {
	Create(ctx context.Context, profile media.Profile) (*media.Profile, error)
	GetByEntityID(ctx context.Context, entityID uuid.UUID) (*media.Profile, error)
	GetByID(ctx context.Context, entityID uuid.UUID) (*media.Profile, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteOthers(ctx context.Context, accountID uuid.UUID) error
}

func NewProfile(logger zerolog.Logger) ProfileDatabase {
	return &Profile{}
}

func (p *Profile) GetByID(ctx context.Context, entityID uuid.UUID) (*media.Profile, error) {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "profile.GetByID").Logger()

	var profile media.Profile

	res := p.db.WithContext(ctx).Model(&media.Profile{}).Where("id = ?", entityID).First(&profile)
	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to get profile")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, sdkerror.ErrRecordNotFound
		}

		return nil, sdkerror.ErrRecordCreation // todo: update accordingly
	}

	return &profile, nil
}

func (p *Profile) GetByEntityID(ctx context.Context, entityID uuid.UUID) (*media.Profile, error) {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "profile.GetByEntityID").Logger()

	var profile media.Profile

	res := p.db.WithContext(ctx).Model(&media.Profile{}).Where("entity_id = ?", entityID).First(&profile)
	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to get profile")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, sdkerror.ErrRecordNotFound
		}

		return nil, sdkerror.ErrRecordCreation // todo: update accordingly
	}

	return &profile, nil
}

func (p *Profile) Create(ctx context.Context, profile media.Profile) (*media.Profile, error) {
	// soft delete others
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "profile.Create").Logger()

	// delete others before trying to create a new one
	//_ = p.DeleteOthers(ctx, profile.AccountID)

	res := p.db.WithContext(ctx).Model(&media.Profile{}).Create(&profile)

	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to create profile")

		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return nil, sdkerror.ErrDuplicateKey
		}

		return nil, sdkerror.ErrRecordCreation
	}

	return &profile, nil
}

func (p *Profile) Delete(ctx context.Context, id uuid.UUID) error {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "profile.Delete").Logger()

	res := p.db.WithContext(ctx).Model(&media.Profile{}).UpdateColumns(media.Profile{ID: id})

	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to delete profile")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrSomethingWentWrong
	}

	return nil
}

func (p *Profile) DeleteOthers(ctx context.Context, accountID uuid.UUID) error {
	log := p.logger.With().
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Str(constants.MethodStrHelper, "profile.DeleteOthers").Logger()

	n := time.Now()

	res := p.db.WithContext(ctx).
		Model(&media.Profile{AccountID: accountID}).
		UpdateColumns(media.Profile{DeletedAt: &n})

	if res.Error != nil {
		log.Err(res.Error).Msg("Failure: failed to delete others")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToUpdateRecord
	}

	return nil
}
