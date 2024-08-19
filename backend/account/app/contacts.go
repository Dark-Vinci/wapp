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

func (a *App) CreateContact(ctx context.Context, contact account.Contacts) (*account.Contacts, error) {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.CreateContact").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msg("creating a user contact")

	c, err := a.contactStore.Create(ctx, contact)
	if err != nil {
		log.Err(err).Msg("failed to create contact")
		return nil, err
	}

	return c, nil
}

func (a *App) BlockContact(ctx context.Context, userID, contactID uuid.UUID) error {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.BlockContact").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msg("Got a request to block a user contact")

	if err := a.contactStore.BlockContact(ctx, contactID, userID); err != nil {
		log.Err(err).Msg("failed to block contact")
		return err
	}

	return nil
}

func (a *App) UnblockContact(ctx context.Context, userID, contactID uuid.UUID) error {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.UnblockContact").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msg("Got a request to unblock a user contact")

	if err := a.contactStore.UnblockContact(ctx, contactID, userID); err != nil {
		log.Err(err).Msg("failed to unblock contact")
		return err
	}

	return nil
}

func (a *App) GetUserContacts(ctx context.Context, contactID uuid.UUID) ([]account.Contacts, error) {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.GetUserContacts").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msg("Got a request to get all users contacts")

	contacts, err := a.contactStore.GetUserContacts(ctx, contactID)
	if err != nil {
		log.Err(err).Msg("failed to get contacts")
		return nil, err
	}

	return contacts, nil
}

func (a *App) GetBlockedContacts(ctx context.Context, userID uuid.UUID) ([]account.Contacts, error) {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.GetBlockedContacts").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msg("Got a request to get blocked contacts")

	contacts, err := a.contactStore.GetBlocked(ctx, userID)
	if err != nil {
		log.Err(err).Msg("failed to get blocked contacts")
		return nil, err
	}

	return contacts, nil
}

func (a *App) RemoveFavouriteContact(ctx context.Context, contactID, userID uuid.UUID) error {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.RemoveContactFromFavourite").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msg("Got a request to remove contact from favourite")

	if err := a.contactStore.MakeUnFavourite(ctx, contactID, userID); err != nil {
		log.Err(err).Msg("failed to remove contact from favourite")
		return err
	}

	return nil
}

func (a *App) MakeContactFavourite(ctx context.Context, contactID, userID uuid.UUID) error {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.MakeContactFavourite").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msg("Got a request to make contact favourite")

	if err := a.contactStore.MakeFavourite(ctx, contactID, userID); err != nil {
		log.Err(err).Msg("Failed to make contact favourite")
		return err
	}

	return nil
}

func (a *App) DeleteContact(ctx context.Context, contactID, userID uuid.UUID) error {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.DeleteContact").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msg("Got a request to delete contact")

	if err := a.contactStore.Delete(ctx, time.Now(), contactID, userID); err != nil {
		log.Err(err).Msg("failed to delete contact")
		return err
	}

	return nil
}

func (a *App) deleteAllUserContacts(ctx context.Context, userID uuid.UUID, tx *gorm.DB) error {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.deleteAllUserContacts").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msg("Got a request to delete all user contacts")

	if err := a.contactStore.DeleteAllUserContacts(ctx, userID, time.Now(), tx); err != nil {
		log.Err(err).Msg("failed to delete all user contacts")
		return err
	}

	return nil
}
