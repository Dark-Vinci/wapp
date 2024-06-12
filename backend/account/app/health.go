package app

import (
	"context"
	"fmt"
)

const logMethodKey string = "LOG_METHOD"

func (a *App) Ping(ctx context.Context, message string) string {
	logger := a.logger.With().
		Str(logMethodKey, "app.Ping").
		Str("message", message).Logger()
	
	logger.Info().Msg("Got a ping request")
	
	return fmt.Sprintf("api GRPC server says %v", message)
}
