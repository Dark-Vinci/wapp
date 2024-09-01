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
	pollVoteStore       *store.PollVoteDatabase
	pollStore           *store.PollDatabase
	pollOptionStore     *store.PollOptionDatabase
	settingsStore       *store.SettingsDatabase
	callStore           *store.CallDatabase
	userCallStore       *store.UserCallDatabase
	db                  *gorm.DB
	logger              *zerolog.Logger
}

func New(z *zerolog.Logger, e *env.Environment) Operations {
	logger := z.With().
		Str(constants.PackageStrHelper, packageName).Logger()

	db := connections.NewDBConn(*z, e)
	channelMessageStore := store.NewChannelMessageDatabase(db)
	messageStore := store.NewMessageDatabase(db)
	groupMessageStore := store.NewGroupMessageDatabase(db)
	pollStore := store.NewPollDatabase(db)
	pollVoteStore := store.NewPollVoteDatabase(db)
	pollOptionStore := store.NewPollOptionDatabase(db)
	settingsStore := store.NewSettingsDatabase(db)
	callStore := store.NewCallDatabase(db)
	userCallStore := store.NewUserCallDatabase(db)

	app := &App{
		logger:              &logger,
		db:                  db.Connection,
		channelMessageStore: channelMessageStore,
		messageStore:        messageStore,
		groupMessageStore:   groupMessageStore,
		pollStore:           pollStore,
		pollVoteStore:       pollVoteStore,
		pollOptionStore:     pollOptionStore,
		settingsStore:       settingsStore,
		callStore:           callStore,
		userCallStore:       userCallStore,
	}

	logger.Info().Msg("Application(app) successfully initialized")

	return Operations(app)
}
