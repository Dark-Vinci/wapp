package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"

	"github.com/dark-vinci/wapp/backend/account/app"
	"github.com/dark-vinci/wapp/backend/account/env"
	"github.com/dark-vinci/wapp/backend/account/server"
	"github.com/dark-vinci/wapp/backend/account/store"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/grpc/account"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
	"github.com/dark-vinci/wapp/backend/sdk/utils/clickhouse"
	"github.com/dark-vinci/wapp/backend/sdk/utils/kafka"
	"github.com/dark-vinci/wapp/backend/sdk/utils/redis"
)

const AppName = "account.main"

func main() {
	_ = os.Setenv("TZ", constants.TimeZone)

	e := env.NewEnv()

	//connect to clickhouse for logs and analytics
	click := clickhouse.New(e.ClickHouseDatabase, e.ClickHouseUsername, e.ClickHousePassword)
	file, _ := os.Open("./logger.text")

	logger := zerolog.New(zerolog.MultiLevelWriter(file, click, zerolog.ConsoleWriter{Out: os.Stdout})).
		With().
		Timestamp().
		Logger()

	appLogger := logger.With().Str(constants.AppNameKey, AppName).Logger()

	promExporter, err := prometheus.New()
	if err != nil {
		appLogger.Fatal().Err(err).Msg("failed to initialize prometheus exporter")
	}

	ctx := context.Background()
	traceExporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		appLogger.Fatal().Err(err).Msg("failed to initialize OT LP trace exporter")
	}

	tracerProvide := trace.NewTracerProvider(trace.WithBatcher(traceExporter))
	otel.SetTracerProvider(tracerProvide)

	meterProvider := metric.NewMeterProvider(metric.WithReader(promExporter))
	otel.SetMeterProvider(meterProvider)

	tracer := otel.Tracer(AppName)

	if e.ShouldMigrate {
		err = utils.Migration(context.Background(), &logger, *e.MigrationConfig(), AppName)

		if err != nil {
			appLogger.Fatal().Err(err).Msg("migration failed")
		}
	}

	red := redis.NewRedis(&logger, "", "", "")
	db := store.NewStore(logger, e)
	kafkaReader := kafka.NewReader([]string{}, "", "", "")
	kafkaWriter := kafka.NewWriter("", "")

	a := app.New(&logger, e, db, *kafkaReader, *kafkaWriter, *red)

	// grpc server initialize
	grpcServer := grpc.NewServer()
	account.RegisterAccountServer(grpcServer, server.New(e, appLogger, a, tracer))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", e.AppPort))

	if err != nil {
		appLogger.Fatal().Err(err).Msg("net.Listen failed")
	}

	appLogger.Info().Msgf("app network is up listening on port %s", e.AppPort)

	defer func() {
		//	close all{redis, db, kafka, server} connection
		click.Close()
		_ = file.Close()
		a.Shutdown()
		_ = listener.Close()
	}()

	appLogger.Info().Msg("serving service over GRPC....")

	go func() {
		if err = grpcServer.Serve(listener); err != nil {
			appLogger.Fatal().Err(err).Msg("grpcServer failed to serve")
		} else {
			// sleep for 1 minute and start consuming from kafka when the application starts
			time.Sleep(60 * time.Second)
			a.Consume()
		}
	}()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		appLogger.Info().Msg("Prometheus metrics server running on port 2112")

		if err = http.ListenAndServe(":2112", nil); err != nil {
			appLogger.Fatal().Err(err).Msg("failed to start Prometheus metrics server")
		}
	}()

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
