package app

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/media/env"
	"github.com/dark-vinci/wapp/backend/media/store"
	"github.com/dark-vinci/wapp/backend/sdk/models/media"
)

const packageName = "media.app"

type Operations interface {
	GetUserProfile(ctx context.Context, userID uuid.UUID) (*media.Profile, error)
	CreateUserProfile(ctx context.Context, userID uuid.UUID, URL string) (*media.Profile, error)
}

type App struct {
	db             *gorm.DB
	logger         *zerolog.Logger
	profileStore   store.ProfileDatabase
	postMediaStore store.PostMediaDatabase
	blurStore      store.BlurDatabase
	chatMediaStore store.ChatMediaDatabase
}

func New(z *zerolog.Logger, e *env.Environment) Operations {
	app := &App{}

	return Operations(app)
}
