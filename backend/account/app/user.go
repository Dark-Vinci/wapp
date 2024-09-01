package app

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

func (a *App) CreateUser(ctx context.Context, u models.User) (*models.User, error) {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.Signup").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Got request to create user account")

	user, err := a.userStore.CreateUser(ctx, u)
	if err != nil {
		log.Err(err).Msg("unable to create user")
		return nil, err
	}

	return user, nil
}

func (a *App) DeleteUserAccount(ctx context.Context, userID uuid.UUID) error {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.DeleteUserAccount").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Got request to delete user account")

	err := a.dbConnection.Transaction(func(tx *gorm.DB) error {
		log.Info().Msg("Transaction: starting transaction to delete user account")

		//	delete all user channel
		//	delete all user group
		//	delete all user contact
		//	delete user

		if err := a.userStore.Delete(ctx, userID, time.Now(), tx); err != nil {
			log.Err(err).Msg("Transaction failed at deleting user")
			return err
		}

		if err := a.deleteAllUserChannels(ctx, userID, tx); err != nil {
			log.Err(err).Msg("Transaction failed at deleting user channels")
			return err
		}

		if err := a.deleteAllUserContacts(ctx, userID, tx); err != nil {
			log.Err(err).Msg("Transaction failed at deleting user contacts")
			return err
		}

		if err := a.deleteAllUserGroups(ctx, userID, tx); err != nil {
			log.Err(err).Msg("Transaction failed at deleting user groups")
			return err
		}

		log.Info().Msg("Transaction: deleted all user entities")

		return nil
	})

	if err != nil {
		log.Err(err).Msg("unable to delete user account")
		return err
	}

	return nil
}
