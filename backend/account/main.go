package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"syscall"

	"github.com/dark-vinci/isok"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/dark-vinci/linkedout/backend/account/app"
	"github.com/dark-vinci/linkedout/backend/account/env"
	"github.com/dark-vinci/linkedout/backend/account/server"
	"github.com/dark-vinci/linkedout/backend/sdk/grpc/account"
	"github.com/dark-vinci/linkedout/backend/sdk/utils"
)

const AppName = "account"

func main() {
	_ = os.Setenv("TZ", "Africa/Lagos")

	f := isok.ResultFun1(os.Create("./zero.log"))

	if f.IsErr() {
		panic("unable to create logger file")
	}

	logger := zerolog.New(f.Unwrap()).With().Timestamp().Logger()
	appLogger := logger.With().Str("APP_NAME", AppName).Logger()

	env := env.NewEnv()

	if env.ShouldMigrate {
		err := utils.Migration(context.Background(), &logger, env, AppName)
		panic(err)
	}

	a := app.New()

	// grpc server initialize
	grpcServer := grpc.NewServer()
	account.RegisterAccountServer(grpcServer, server.New(env, appLogger, a))

	res := isok.ResultFun1(net.Listen("tcp", fmt.Sprintf(":%s", env.AppPort)))

	if res.IsErr() {
		appLogger.Fatal().Err(res.UnwrapErr()).Msg("net.Listen failed")
		panic(res.UnwrapErr())
	}

	listener := res.Unwrap()

	appLogger.Info().Msgf("app network is up listening on port %s", env.AppPort)

	defer func() {
		_ = listener.Close()
	}()

	appLogger.Info().Msg("serving service over GRPC....")
	serv := isok.ResultFun0(grpcServer.Serve(listener))

	if serv.IsErr() {
		appLogger.Fatal().Err(serv.UnwrapErr()).Msg("grpcServer failed to serve")
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
