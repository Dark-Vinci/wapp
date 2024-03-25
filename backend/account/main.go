package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"syscall"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/dark-vinci/linkedout/backend/account/app"
	"github.com/dark-vinci/linkedout/backend/account/server"
	"github.com/dark-vinci/linkedout/backend/sdk/grpc/account"
	"github.com/dark-vinci/linkedout/backend/sdk/models"
	"github.com/dark-vinci/linkedout/backend/sdk/utils"
)

const AppName = "account"

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	appLogger := logger.With().Str("APP_NAME", AppName).Logger()

	env := models.Env{}

	if env.ShouldMigrate {
		err := utils.Migration(context.Background(), &logger, &env, AppName)
		panic(err)
	}

	a := app.New()

	// grpc server initialize
	grpcServer := grpc.NewServer()
	account.RegisterAccountServer(grpcServer, server.New(&env, appLogger, a))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", env.AppPort))
	appLogger.Info().Msgf("app network is up listening on port :%s", env.AppPort)

	defer func() {
		_ = listener.Close()
	}()

	if err != nil {
		appLogger.Fatal().Err(err).Msg("net.Listen failed")
	}

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

		utils.ShutdownGracefully(appLogger, grpcServer, 32)
	}
}
