package app

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/dark-vinci/linkedout/backend/account/connection"
	"github.com/dark-vinci/linkedout/backend/account/env"
	"github.com/dark-vinci/linkedout/backend/account/store"
)

// const packageName string = "account.app"

type Operations interface {
	Ping(ctx context.Context, message string) string
	Signup(ctx context.Context, a int) (*int, error)
}

type App struct {
	env       *env.Environment
	red       *connection.RedisClient
	kafka     *connection.Kafka
	userStore *store.UserDatabase
	logger    zerolog.Logger
}

func New(z *zerolog.Logger, e *env.Environment) Operations {
	red := connection.NewRedisClient(z, e)
	db := connection.NewDBConn(*z, e)
	uStore := store.NewUser(db)
	kafka := connection.NewKafka(*z, e)

	app := &App{
		red:    red,
		env:    e,
		logger: *z,
		userStore: uStore,
		kafka: kafka,
	}

	return Operations(app)
}
