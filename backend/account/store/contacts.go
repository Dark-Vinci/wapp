package store

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/account/connection"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models/account"
	"github.com/dark-vinci/wapp/backend/sdk/sdkerror"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

type Contact struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

type ContactDatabase interface {
	BlockContact(ctx context.Context, contactID, userID uuid.UUID) error
	GetUserContacts(ctx context.Context, userID uuid.UUID) ([]account.Contacts, error)
	Create(ctx context.Context, contact account.Contacts) (*account.Contacts, error)
	Delete(ctx context.Context, deletedAt time.Time, contactID, userID uuid.UUID) error
	UnblockContact(ctx context.Context, contactID, userID uuid.UUID) error
	MakeFavourite(ctx context.Context, contactID uuid.UUID, userID uuid.UUID) error
	MakeUnFavourite(ctx context.Context, contactID uuid.UUID, userID uuid.UUID) error
	GetBlocked(ctx context.Context, userID uuid.UUID) ([]account.Contacts, error)
	DeleteAllUserContacts(ctx context.Context, userID uuid.UUID, deletedAt time.Time, tx *gorm.DB) error
}

func NewContact(conn *connection.DBConn) *ContactDatabase {
	l := conn.Log.With().
		Str(constants.FunctionNameHelper, "NewContact").
		Str(constants.PackageStrHelper, packageName).
		Logger()

	contact := &Contact{
		logger: &l,
		db:     conn.Connection,
	}

	operations := ContactDatabase(contact)

	return &operations
}

func (c *Contact) Create(ctx context.Context, contact account.Contacts) (*account.Contacts, error) {
	log := c.logger.With().
		Str(constants.MethodStrHelper, "contact.Create").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Got a request to create a contact with info %v", contact)

	if res := c.db.WithContext(ctx).Model(&account.Contacts{}).Create(&contact); res.Error != nil {
		log.Err(res.Error).Msg("Failed to create contact")

		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return nil, sdkerror.ErrDuplicateKey
		}

		return nil, sdkerror.ErrRecordCreation
	}

	return &contact, nil
}

func (c *Contact) MakeFavourite(ctx context.Context, contactID uuid.UUID, userID uuid.UUID) error {
	log := c.logger.With().
		Str(constants.MethodStrHelper, "contact.MakeFavourite").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Got a request to make a contact with info %v a fourite of %v", contactID, userID)

	if res := c.db.WithContext(ctx).Model(&account.Contacts{}).Where("id = ? and owner_id = ?", contactID, userID).UpdateColumns(&account.Contacts{IsFavorite: true}); res.Error != nil {
		log.Err(res.Error).Msg("Failed to make contact a favourite")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToUpdateRecord
	}

	return nil
}

func (c *Contact) MakeUnFavourite(ctx context.Context, contactID uuid.UUID, userID uuid.UUID) error {
	log := c.logger.With().
		Str(constants.MethodStrHelper, "contact.MakeUnFavourite").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Got a request to make a contact with info %v a unfourite of %v", contactID, userID)

	if res := c.db.WithContext(ctx).Model(&account.Contacts{}).UpdateColumns(&account.Contacts{IsFavorite: false}); res.Error != nil {
		log.Err(res.Error).Msg("Failed to make contact a un favourite")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToUpdateRecord
	}

	return nil
}

func (c *Contact) GetBlocked(ctx context.Context, userID uuid.UUID) ([]account.Contacts, error) {
	log := c.logger.With().
		Str(constants.MethodStrHelper, "contact.GetBlocked").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Got a request to get a blocked contacts with info %v", userID)

	var result []account.Contacts

	if err := c.db.WithContext(ctx).Model(&account.Contacts{}).Where("is_blocked = true AND owner_id = ?", userID).Find(&result).Error; err != nil {
		log.Err(err).Msg("Failed to get blocked contacts")

		return nil, sdkerror.ErrRecordNotFound
	}

	return result, nil
}

func (c *Contact) GetUserContacts(ctx context.Context, userID uuid.UUID) ([]account.Contacts, error) {
	log := c.logger.With().
		Str(constants.MethodStrHelper, "contact.GetUserContacts").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Got a request to get contacts for a user with id %v", userID)

	var contacts []account.Contacts

	if res := c.db.WithContext(ctx).Model(&account.Contacts{}).Where("owner_id = ?", userID).Find(&contacts); res.Error != nil {
		log.Err(res.Error).Msg("Failed to get contacts")

		return nil, sdkerror.ErrRecordNotFound
	}

	return contacts, nil
}

func (c *Contact) BlockContact(ctx context.Context, contactID, userID uuid.UUID) error {
	log := c.logger.With().
		Str(constants.MethodStrHelper, "contact.BlockContact").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Got a request to block a contact with id %s with user %s", contactID.String(), userID.String())

	if res := c.db.WithContext(ctx).Model(&account.Contacts{}).Where("owner_id = ? AND id = ?", userID, contactID).UpdateColumns(account.Contacts{IsBlocked: true}); res.Error != nil {
		log.Err(res.Error).Msg("Failed to block contact")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrRecordCreation // UPDATE
	}

	return nil
}

func (c *Contact) UnblockContact(ctx context.Context, contactID, userID uuid.UUID) error {
	log := c.logger.With().
		Str(constants.MethodStrHelper, "contact.UnblockContact").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Got a request to unblock a contact with id %s with user %s", contactID.String(), userID.String())

	if res := c.db.WithContext(ctx).Model(&account.Contacts{}).Where("owner_id = ? AND id = ?", userID, contactID).UpdateColumns(account.Contacts{IsBlocked: false}); res.Error != nil {
		log.Err(res.Error).Msg("Failed to unblock contact")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToUpdateRecord
	}

	return nil
}

func (c *Contact) Delete(ctx context.Context, deletedAt time.Time, contactID, userID uuid.UUID) error {
	log := c.logger.With().
		Str(constants.MethodStrHelper, "contact.Delete").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Got a request to delete a contact with id %s with user %s", contactID.String(), userID.String())

	if res := c.db.WithContext(ctx).Model(&account.Contacts{}).Where("owner_id = ? AND id = ?", userID, contactID).UpdateColumns(account.Contacts{DeletedAt: deletedAt}); res.Error != nil {
		log.Err(res.Error).Msg("Failed to delete contact")

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToDeleteRecord
	}

	return nil
}

func (c *Contact) DeleteAllUserContacts(ctx context.Context, userID uuid.UUID, deletedAt time.Time, tx *gorm.DB) error {
	log := c.logger.With().
		Str(constants.MethodStrHelper, "contact.DeleteAllUserContacts").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Got a request to delete all contacts for a user with id %s", userID)

	if tx != nil {
		log.Info().Msg("deleting all user contact in a transaction")

		if err := tx.WithContext(ctx).Model(&account.Contacts{}).Where("user_id = ?", userID).UpdateColumns(&account.Contacts{DeletedAt: deletedAt}).Error; err != nil {
			log.Err(err).Msg("Failed to delete all contacts")

			if errors.Is(err, gorm.ErrRecordNotFound) {
				return sdkerror.ErrRecordNotFound
			}

			return sdkerror.ErrFailedToDeleteRecord
		}
	}

	return nil
}
