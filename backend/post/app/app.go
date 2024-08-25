package app

import (
	"github.com/dark-vinci/wapp/backend/post/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

const packageName = "post.app"

type App struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

type Operations interface{}

func New(z *zerolog.Logger, e *env.Environment) Operations {
	logger := z.With().
		Str(constants.FunctionNameHelper, "app.New").
		Str(constants.PackageStrHelper, packageName).Logger()

	app := &App{
		logger: &logger,
	}

	return Operations(app)
}
