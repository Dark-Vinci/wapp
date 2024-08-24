package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"syscall"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/dark-vinci/wapp/backend/chats/app"
	"github.com/dark-vinci/wapp/backend/chats/env"
	"github.com/dark-vinci/wapp/backend/chats/server"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/grpc/chat"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
	"github.com/dark-vinci/wapp/backend/sdk/utils/clickhouse"
)

const AppName = "chat.main"
const AppNameKey = "APP_NAME"

func main() {
	_ = os.Setenv("TZ", constants.TimeZone)

	e := env.NewEnv()

	//connect to clickhouse for logs and analytics
	click := clickhouse.New(e.ClickHouseDatabase, e.ClickHouseUsername, e.ClickHousePassword)

	defer click.Close()

	logger := zerolog.New(click).With().Timestamp().Logger()
	appLogger := logger.With().Str(AppNameKey, AppName).Logger()

	if e.ShouldMigrate {
		err := utils.Migration(context.Background(), &logger, *e.MigrationConfig(), AppName)

		if err != nil {
			appLogger.Fatal().Err(err).Msg("migration failed")
			panic(err)
		}
	}

	a := app.New(&logger, e)

	// grpc server initialize
	grpcServer := grpc.NewServer()
	chat.RegisterAccountServer(grpcServer, server.New(e, appLogger, a))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", e.AppPort))

	if err != nil {
		appLogger.Fatal().Err(err).Msg("net.Listen failed")
		panic(err)
	}

	appLogger.Info().Msgf("app network is up listening on port %s", e.AppPort)

	defer func() {
		_ = listener.Close()
	}()

	appLogger.Info().Msg("serving service over GRPC....")

	if err = grpcServer.Serve(listener); err != nil {
		appLogger.Fatal().Err(err).Msg("grpcServer failed to serve")
		panic("unable to start service at this time")
	}

	// initialize shutdown handling
	defer utils.RecoverAndLogPanic(appLogger)
	shutdown := utils.ListenForShutdown()
	serverErrors := make(chan error, 1)

	select {
	case err = <-serverErrors:
		appLogger.Panic().Err(err).Msg("server error")
	case sig := <-shutdown:
		appLogger.Info().Msgf("%v : start server shutdown.", sig)

		if sig == syscall.SIGSTOP {
			appLogger.Info().Msg("integrity issue caused shutdown")
		}

		utils.ShutdownGracefully(appLogger, grpcServer, nil)
	}
}
