package app

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/post/connections"
	"github.com/dark-vinci/wapp/backend/post/env"
	"github.com/dark-vinci/wapp/backend/post/store"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

const packageName = "post.app"

type App struct {
	logger           *zerolog.Logger
	db               *gorm.DB
	contactPostStore *store.ContactPostDatabase
	postStore        *store.PostDatabase
	postViewStore    *store.PostViewDatabase
	settingsStore    *store.SettingsDatabase
}

type Operations interface{}

func New(z *zerolog.Logger, e *env.Environment) Operations {
	logger := z.With().
		Str(constants.FunctionNameHelper, "app.New").
		Str(constants.PackageStrHelper, packageName).Logger()

	dbConn := connections.NewDBConn(*z, e)
	contactPostStore := store.NewContactPost(dbConn)
	postStore := store.NewPost(dbConn)
	postViewStore := store.NewPostView(dbConn)
	settingsStore := store.NewSettings(dbConn)

	app := &App{
		logger:           &logger,
		db:               dbConn.Connection,
		contactPostStore: contactPostStore,
		postStore:        postStore,
		postViewStore:    postViewStore,
		settingsStore:    settingsStore,
	}

	return Operations(app)
}
