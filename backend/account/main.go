package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/dark-vinci/linkedout/backend/account/app"
	"github.com/dark-vinci/linkedout/backend/account/server"
	"github.com/dark-vinci/linkedout/backend/sdk/grpc/account"
	"github.com/dark-vinci/linkedout/backend/sdk/models"
	"github.com/dark-vinci/linkedout/backend/sdk/utils"
)

const AppName = "account"

type DataSaver struct {
	Data   []string
	DBChan chan string // Channel to send data to be written to the database
}

// Write appends the data to the Data slice and sends it to the DBChan for database write
func (d *DataSaver) Write(p []byte) (int, error) {
	// Convert byte slice to string and append it to Data slice
	str := string(p)
	d.Data = append(d.Data, str)

	// Send the data to be written to the database through a channel
	go func() {
		d.DBChan <- str
	}()

	// Return the number of bytes written and no error
	return len(p), nil
}

// Function to simulate database write
func writeToDB(dbChan <-chan string) {
	for {
		data, ok := <-dbChan
		if !ok {
			return // Channel closed
		}
		// Simulate writing to the database
		fmt.Println("Writing to database:", data)
		time.Sleep(1 * time.Second) // Simulating database write time
	}
}

func main() {
	_ = os.Setenv("TZ", "Africa/Lagos")

	//zlFile, err := os.Create("./zero.log")
	//if err != nil {
	//	panic("cant create file")
	//}

	myDataSaver := &DataSaver{
		DBChan: make(chan string),
	}

	// Start database write Goroutine
	go writeToDB(myDataSaver.DBChan)

	logger := zerolog.New(myDataSaver).With().Timestamp().Logger()
	appLogger := logger.With().Str("APP_NAME", AppName).Logger()

	env := models.Env{
		AppPort: "8080",
	}

	if env.ShouldMigrate {
		err := utils.Migration(context.Background(), &logger, &env, AppName)
		panic(err)
	}

	a := app.New()

	// grpc server initialize
	grpcServer := grpc.NewServer()
	account.RegisterAccountServer(grpcServer, server.New(&env, appLogger, a))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", env.AppPort))
	appLogger.Info().Msgf("app network is up listening on port %s", env.AppPort)

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

		utils.ShutdownGracefully(appLogger, grpcServer, nil)
	}
}
