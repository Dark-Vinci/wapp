package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"

	"github.com/dark-vinci/wapp/backend/gateway/env"
	"github.com/dark-vinci/wapp/backend/gateway/handlers"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/utils/clickhouse"
)

func main() {
	_ = os.Setenv("TZ", constants.TimeZone)

	e := env.New()

	click := clickhouse.New("", "", "")

	logger := zerolog.New(click).With().Timestamp().Logger()
	appLogger := logger.With().Str("GATEWAY", "api").Logger()

	appLogger.Debug().Msg("something should happen")
	appLogger.Debug().Msg("another log in the logger file")

	h := handlers.New(e, logger)
	h.Build()

	server := &http.Server{
		Addr:    ":8080",
		Handler: h.GetEngine(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			appLogger.Fatal().Err(err).Msg("failed to start server")
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	<-ctx.Done()

	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		appLogger.Err(err).Msg("Server forced to shutdown")
	}

	appLogger.Debug().Msg("server last message")
}
