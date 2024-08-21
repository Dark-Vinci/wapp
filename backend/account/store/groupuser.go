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

type GroupUser struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

type GroupUserDatabase interface {
	Create(ctx context.Context, groupUser account.GroupUser) (*account.GroupUser, error)
	RemoveUser(ctx context.Context, deletedAt time.Time, userID, groupID uuid.UUID) error
	UnMuteGroup(ctx context.Context, groupID, userID uuid.UUID) error
	MuteGroup(ctx context.Context, groupID, userID uuid.UUID) error
	FindAllGroupUser(ctx context.Context, groupID uuid.UUID) ([]account.GroupUser, error)
	FindUserByID(ctx context.Context, userID, groupID uuid.UUID) (*account.GroupUser, error)
}

func NewGroupUser(conn *connection.DBConn) *GroupUserDatabase {
	logger := conn.Log.With().
		Str(constants.FunctionNameHelper, "NewGroupUserStore").
		Str(constants.PackageStrHelper, packageName).
		Logger()

	groupUser := &GroupUser{
		logger: &logger,
		db:     conn.Connection,
	}

	operations := GroupUserDatabase(groupUser)

	return &operations
}

func (gu *GroupUser) Create(ctx context.Context, groupUser account.GroupUser) (*account.GroupUser, error) {
	log := gu.logger.With().
		Str(constants.MethodStrHelper, "groupUser.Create").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msg("Got a request to Create a group's user")

	if err := gu.db.WithContext(ctx).Model(&account.GroupUser{}).Create(&groupUser).Error; err != nil {
		log.Err(err).Msg("Failed to create group user")

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, sdkerror.ErrDuplicateKey
		}

		return nil, sdkerror.ErrRecordCreation
	}

	return &groupUser, nil
}

func (gu *GroupUser) RemoveUser(ctx context.Context, deletedAt time.Time, userID, groupID uuid.UUID) error {
	log := gu.logger.With().
		Str(constants.MethodStrHelper, "groupUser.RemoveUser").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msg("Got a request to Remove a user")

	if err := gu.db.WithContext(ctx).Model(&account.GroupUser{}).Where("user_id = ? and group_id = ?", userID, groupID).UpdateColumns(&account.GroupUser{DeletedAt: &deletedAt}).Error; err != nil {
		log.Err(err).Msg("Failed to remove a user")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToUpdateRecord
	}

	return nil
}

func (gu *GroupUser) UnMuteGroup(ctx context.Context, groupID, userID uuid.UUID) error {
	log := gu.logger.With().
		Str(constants.MethodStrHelper, "groupUser.UnMuteGroup").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msg("Got a request to mute a group")

	if err := gu.db.WithContext(ctx).Model(&account.GroupUser{}).Where("user_id = ? and group_id = ?", userID, groupID).UpdateColumns(&account.GroupUser{Mute: false}).Error; err != nil {
		log.Err(err).Msg("Failed to unmute a group")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToUpdateRecord
	}

	return nil
}

func (gu *GroupUser) MuteGroup(ctx context.Context, groupID, userID uuid.UUID) error {
	log := gu.logger.With().
		Str(constants.MethodStrHelper, "groupUser.MuteGroup").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msg("Got a request to mute a group")

	if err := gu.db.WithContext(ctx).Model(&account.GroupUser{}).Where("user_id = ? and group_id = ?", userID, groupID).UpdateColumns(&account.GroupUser{Mute: true}).Error; err != nil {
		log.Err(err).Msg("Failed to mute a group")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrFailedToUpdateRecord
	}

	return nil
}

func (gu *GroupUser) FindAllGroupUser(ctx context.Context, groupID uuid.UUID) ([]account.GroupUser, error) {
	log := gu.logger.With().
		Str(constants.MethodStrHelper, "groupUser.FindAllGroupUser").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msg("Got a request to Find all group users")

	var groupUsers []account.GroupUser

	if err := gu.db.WithContext(ctx).Model(&account.GroupUser{}).Where("group_id = ?", groupID).Find(&groupUsers).Error; err != nil {
		log.Err(err).Msg("Failed to find all group users")

		return nil, sdkerror.ErrRecordNotFound
	}

	return groupUsers, nil
}

func (gu *GroupUser) FindUserByID(ctx context.Context, userID, groupID uuid.UUID) (*account.GroupUser, error) {
	log := gu.logger.With().
		Str(constants.MethodStrHelper, "groupUser.FindUserByID").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msg("Got a request to Find a user")

	var result account.GroupUser

	if err := gu.db.WithContext(ctx).Model(&account.GroupUser{}).Where("user_id = ? and group_id = ?", userID, groupID).First(&result).Error; err != nil {
		log.Err(err).Msg("Failed to find a user")
		return nil, sdkerror.ErrRecordNotFound
	}

	return &result, nil
}
