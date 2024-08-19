package app

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models/account"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

func (a *App) CreateGroup(ctx context.Context, group account.Group) error {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.CreateGroup").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Got a request to create group")

	if _, err := a.groupStore.Create(ctx, group); err != nil {
		log.Err(err).Msg("failed to create group")
		return err
	}

	return nil
}

func (a *App) DeleteGroup(ctx context.Context, groupID uuid.UUID) error {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.DeleteGroup").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Got a request to delete group")

	if err := a.groupStore.DeleteByID(ctx, groupID, time.Now()); err != nil {
		log.Err(err).Msg("failed to delete group")
		return err
	}

	return nil
}

func (a *App) deleteAllUserGroups(ctx context.Context, userID uuid.UUID, tx *gorm.DB) error {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.DeleteAllUserGroups").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Got a request to delete all user groups")

	if err := a.groupStore.DeleteAllUserGroup(ctx, userID, time.Now(), tx); err != nil {
		log.Err(err).Msg("failed to delete all user groups")
		return err
	}

	return nil
}
