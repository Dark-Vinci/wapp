package app

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/chats/connections"
	"github.com/dark-vinci/wapp/backend/chats/env"
	"github.com/dark-vinci/wapp/backend/chats/store"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

const packageName = "chats.app"

type Operations interface{}

type App struct {
	channelMessageStore *store.ChannelMessageDatabase
	messageStore        *store.MessageDatabase
	groupMessageStore   *store.GroupMessageDatabase
	db                  *gorm.DB
	logger              *zerolog.Logger
}

func New(z *zerolog.Logger, e *env.Environment) Operations {
	logger := z.With().
		Str(constants.FunctionNameHelper, "New").
		Str(constants.PackageStrHelper, packageName).
		Logger()

	db := connections.NewDBConn(*z, e)
	channelMessageStore := store.NewChannelMessageDatabase(db)
	messageStore := store.NewMessageDatabase(db)
	groupMessageStore := store.NewGroupMessageDatabase(db)

	app := &App{
		logger:              &logger,
		db:                  db.Connection,
		channelMessageStore: channelMessageStore,
		messageStore:        messageStore,
		groupMessageStore:   groupMessageStore,
	}

	logger.Info().Msg("Application(app) successfully initialized")

	return Operations(app)
}
