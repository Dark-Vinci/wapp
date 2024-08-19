package app

import (
	"context"
	"fmt"

	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

func (a *App) Ping(ctx context.Context, message string) string {
	logger := a.logger.With().
		Str(constants.MethodStrHelper, "app.Ping").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	logger.Info().Msg("Got a ping request")

	return fmt.Sprintf("api GRPC server says %v", message)
}
