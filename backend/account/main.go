package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"syscall"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/dark-vinci/wapp/backend/account/app"
	"github.com/dark-vinci/wapp/backend/account/env"
	"github.com/dark-vinci/wapp/backend/account/server"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/grpc/account"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

const AppName = "account"

func main() {
	_ = os.Setenv("TZ", constants.TimeZone)

	f, err := os.Create("./zero.log")

	if err != nil {
		panic("unable to create logger file")
	}

	logger := zerolog.New(f).With().Timestamp().Logger()
	appLogger := logger.With().Str("APP_NAME", AppName).Logger()

	e := env.NewEnv()

	if e.ShouldMigrate {
		err := utils.Migration(context.Background(), &logger, *e.MigrationConfig(), AppName)
		panic(err)
	}

	a := app.New(&logger, e)

	// grpc server initialize
	grpcServer := grpc.NewServer()
	account.RegisterAccountServer(grpcServer, server.New(e, appLogger, a))

	res, err := net.Listen("tcp", fmt.Sprintf(":%s", e.AppPort))

	if err != nil {
		appLogger.Fatal().Err(err).Msg("net.Listen failed")
		panic(err)
	}

	listener := res

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
	case err := <-serverErrors:
		appLogger.Panic().Err(err).Msg("server error")
	case sig := <-shutdown:
		appLogger.Info().Msgf("%v : start server shutdown.", sig)

		if sig == syscall.SIGSTOP {
			appLogger.Info().Msg("integrity issue caused shutdown")
		}

		utils.ShutdownGracefully(appLogger, grpcServer, nil)
	}
}
