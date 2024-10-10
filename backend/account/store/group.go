package store

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models"
	"github.com/dark-vinci/wapp/backend/sdk/models/account"
	"github.com/dark-vinci/wapp/backend/sdk/sdkerror"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

type Group struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source group.go -destination ./mock/group_mock.go -package mock GroupDatabase
type GroupDatabase interface {
	Create(ctx context.Context, group account.Group) (*account.Group, error)
	GetGroupByID(ctx context.Context, id uuid.UUID) (*models.Group, error)
	DeleteByID(ctx context.Context, id uuid.UUID, deletedAt time.Time) error
	DeleteAllUserGroup(ctx context.Context, userID uuid.UUID, deletedAt time.Time, tx *gorm.DB) error
}

func NewGroup(conn *Store) *GroupDatabase {
	l := conn.Log.With().
		Str(constants.FunctionNameHelper, "NewGroup").
		Str(constants.PackageStrHelper, packageName).
		Logger()

	group := &Group{
		logger: &l,
		db:     conn.Connection,
	}

	operations := GroupDatabase(group)
	return &operations
}

func (g *Group) Create(ctx context.Context, group account.Group) (*account.Group, error) {
	log := g.logger.With().
		Str(constants.MethodStrHelper, "group.Create").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Creating a new group")

	res := g.db.WithContext(ctx).Model(&models.Group{}).Create(&group)

	if res.Error != nil {
		log.Err(res.Error).Msg("failed to create group")

		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return nil, sdkerror.ErrDuplicateKey
		}

		return nil, sdkerror.ErrRecordCreation
	}

	return &group, nil
}

func (g *Group) GetGroupByID(ctx context.Context, id uuid.UUID) (*models.Group, error) {
	log := g.logger.With().
		Str(constants.MethodStrHelper, "group.GetGroupByID").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("Got a request to get a group by ID")

	var group models.Group
	res := g.db.WithContext(ctx).Model(&models.Group{}).Where("id = ?", id).First(&group)

	if res.Error != nil {
		log.Err(res.Error).Msg("failed to get the group by ID")

		return nil, sdkerror.ErrRecordNotFound
	}

	return &group, nil
}

func (g *Group) DeleteAllUserGroup(ctx context.Context, userID uuid.UUID, deletedAt time.Time, tx *gorm.DB) error {
	log := g.logger.With().
		Str(constants.MethodStrHelper, "group.DeleteAllUserGroup").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Deleting all user with id: %v groups", userID)

	if tx != nil {
		if err := tx.WithContext(ctx).Model(&account.Group{}).Where("owner_id = ?", userID).UpdateColumns(&account.Group{DeletedAt: &deletedAt}).Error; err != nil {
			log.Err(err).Msg("failed to delete user groups")
			return err
		}

		return nil
	}

	return nil
}

func (g *Group) DeleteByID(ctx context.Context, id uuid.UUID, deletedAt time.Time) error {
	log := g.logger.With().
		Str(constants.MethodStrHelper, "group.DeleteByID").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msgf("Got a request to delete a group %s by ID", id.String())

	res := g.db.WithContext(ctx).Model(&models.Group{}).Where("id = ?", id).UpdateColumns(models.Group{DeletedAt: &deletedAt})

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToDeleteRecord
	}

	return nil
}
