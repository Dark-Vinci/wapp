package app

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models/media"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

func (a *App) GetProfile(ctx context.Context, entityID uuid.UUID) (*media.Profile, error) {
	logger := a.logger.With().
		Str(constants.MethodStrHelper, "app.GetUserProfile").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	profile, err := a.profileStore.GetByID(ctx, entityID)
	if err != nil {
		logger.Err(err).Msg("failed to delete user")
		return nil, err
	}

	return profile, nil
}

func (a *App) CreateGroupProfile(ctx context.Context, entityID uuid.UUID, userID uuid.UUID, URL string) (*media.Profile, error) {
	logger := a.logger.With().
		Str(constants.MethodStrHelper, "app.CreateGroupProfile").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	if err := a.profileStore.DeleteOthers(ctx, entityID); err != nil {
		logger.Err(err).Msg("failed to delete other entity group profile")
		return nil, err
	}

	nProfile := media.Profile{
		URL:       URL,
		AccountID: entityID,
		CreatedBy: userID,
		CreatedAt: time.Now(),
	}

	profile, err := a.profileStore.Create(ctx, nProfile)
	if err != nil {
		logger.Err(err).Msg("failed to create group")
		return nil, err
	}

	return profile, nil
}

func (a *App) CreateUserProfile(ctx context.Context, userID uuid.UUID, URL string) (*media.Profile, error) {
	logger := a.logger.With().
		Str(constants.MethodStrHelper, "app.CreateUserProfile").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	newProfile := media.Profile{
		AccountID: userID,
		URL:       URL,
		CreatedBy: userID,
		CreatedAt: time.Now(),
	}

	// delete others first
	err := a.profileStore.DeleteOthers(ctx, userID)
	if err != nil {
		logger.Err(err).Msg("failed to delete user")
		return nil, err
	}

	profile, err := a.profileStore.Create(ctx, newProfile)
	if err != nil {
		logger.Err(err).Msg("AppError: failure creating user profile")
		return nil, err
	}

	return profile, nil
}
