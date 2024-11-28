package app

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/media/env"
	"github.com/dark-vinci/wapp/backend/media/store"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models/media"
)

const packageName = "media.app"

type Operations interface {
	GetProfile(ctx context.Context, userID uuid.UUID) (*media.Profile, error)
	CreateUserProfile(ctx context.Context, userID uuid.UUID, URL string) (*media.Profile, error)
	CreateGroupProfile(ctx context.Context, entityID uuid.UUID, userID uuid.UUID, URL string) (*media.Profile, error)
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
	st := store.New(*z, e)
	profileStore := store.NewProfile(*z, st.Connection)
	postMediaStore := store.NewPostMedia(*z, st.Connection)
	blurStore := store.NewBlur(*z, st.Connection)
	chatMediaStore := store.NewChatMedia(*z, st.Connection)

	log := z.With().Str(constants.PackageStrHelper, packageName).Logger()

	app := &App{
		logger:         &log,
		db:             st.Connection,
		profileStore:   profileStore,
		chatMediaStore: chatMediaStore,
		blurStore:      blurStore,
		postMediaStore: postMediaStore,
	}

	return Operations(app)
}
