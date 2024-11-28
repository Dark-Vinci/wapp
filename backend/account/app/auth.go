package app

import (
	"context"

	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models/account"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *App) LoginToAccount(ctx context.Context, username, password string) error {
	logger := a.logger.With().
		Str(constants.MethodStrHelper, "app.LoginToAccount").
		Str(constants.RequestID, utils.GetRequestID(ctx)).
		Logger()

	logger.Info().Msg("Got a request to log in user")

	user, err := a.userStore.

	return nil
}

func (a *App) Logout(ctx context.Context) error {
	return nil
}

func (a *App) CreateAccount(ctx context.Context, details LoginRequest) (*account.User, error) {
	return nil, nil
}

func (a *App) VerifyOTP(ctx context.Context, otp string) error {
	return nil
}
