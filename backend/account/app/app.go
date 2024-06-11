package app

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/dark-vinci/linkedout/backend/account/connection"
	"github.com/dark-vinci/linkedout/backend/account/env"
)

const packageName string = "account.app"

type Operations interface {
	Signup(ctx context.Context, a int) (*int, error)
}

type App struct {
	env    *env.Environment
	red    *connection.RedisClient
	kafka  *connection.Kafka
	logger zerolog.Logger
}

func New(z *zerolog.Logger, e *env.Environment) Operations {
	red := connection.NewRedisClient(z, e)

	app := &App{
		red:    red,
		env:    e,
		logger: *z,
	}

	return Operations(app)
}
