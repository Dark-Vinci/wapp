package app

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/media/env"
)

const packageName = "media.app"

type Operations interface{}

type App struct {
	db *gorm.DB
}

func New(z *zerolog.Logger, e *env.Environment) Operations {
	app := &App{}

	return Operations(app)
}
