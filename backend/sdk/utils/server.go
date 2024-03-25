package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func ListenForShutdown() chan os.Signal {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	return shutdown
}

func RecoverAndLogPanic(log zerolog.Logger) {
	if err := recover(); err != nil {
		log.Panic().Str("crash", fmt.Sprintf("%+v", err)).Stack().Msg("recovered in main")
	}
}

// ShutdownGracefully close all connections properly
func ShutdownGracefully(logger zerolog.Logger, grpcServer *grpc.Server, storage interface{}) {
	logger.Info().Msg("shutting down app gracefully...")
	if grpcServer != nil {
		grpcServer.GracefulStop()
		logger.Info().Msg("grpc server shutdown completed gracefully")
	} else {
		logger.Info().Msg("grpc server wasn't live before")
	}

	if storage != nil {
		defer func() {
			//storage.Close()
			logger.Info().Msg("postgres client disconnected gracefully")
		}()
	} else {
		logger.Info().Msg("postgres client wasn't connected before")
	}
}
