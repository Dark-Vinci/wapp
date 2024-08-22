package app

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models/account"
	"github.com/dark-vinci/wapp/backend/sdk/sdkerror"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

func (a *App) CreateChannel(ctx context.Context, channel account.Channel) (*account.Channel, error) {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.CreateChannel").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	ch, err := a.channelStore.Create(ctx, channel)
	if err != nil {
		log.Err(err).Msg("failed to create channel")
		return nil, err
	}

	return ch, nil
}

func (a *App) deleteAllUserChannels(ctx context.Context, userID uuid.UUID, tx *gorm.DB) error {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.DeleteAllUserChannels").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	if err := a.channelStore.DeleteUserChannels(ctx, userID, time.Now(), tx); err != nil {
		log.Err(err).Msg("failed to delete user channels")
		return err
	}

	return nil
}

func (a *App) AddUserToChannel(ctx context.Context, userID, channelID uuid.UUID) error {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.AddUserToChannel").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	log.Info().Msg("adding user to channel")

	if _, err := a.channelStore.GetChannelByID(ctx, channelID); err != nil {
		log.Err(err).Msg("failed to get channel by id")

		if errors.Is(err, sdkerror.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrSomethingWentWrong
	}

	if _, err := a.userStore.GetUserByID(ctx, userID); err != nil {
		log.Err(err).Msg("failed get user by id")
		if errors.Is(err, sdkerror.ErrRecordNotFound) {
			return sdkerror.ErrRecordNotFound
		}

		return sdkerror.ErrSomethingWentWrong
	}

	channelUser := account.ChannelUser{
		ChannelID: channelID,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := a.channelUserStore.Create(ctx, channelUser); err != nil {
		log.Err(err).Msg("failed add user channel to channel")
		return err
	}

	return nil
}

func (a *App) DeleteChannel(ctx context.Context, channelID uuid.UUID) error {
	log := a.logger.With().
		Str(constants.MethodStrHelper, "app.DeleteChannel").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	if err := a.channelStore.DeleteByID(ctx, channelID, time.Now()); err != nil {
		log.Err(err).Msg("failed to delete channel")
		return err
	}

	return nil
}
