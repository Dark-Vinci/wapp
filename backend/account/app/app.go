package app

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/account/connection"
	"github.com/dark-vinci/wapp/backend/account/env"
	"github.com/dark-vinci/wapp/backend/account/store"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models"
	"github.com/dark-vinci/wapp/backend/sdk/models/account"
)

const packageName string = "account.app"

//go:generate mockgen -source app.go -destination ./mock/mock_app.go -package mock  Operations
type Operations interface {
	DeleteUserAccount(ctx context.Context, userID uuid.UUID) error
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	CreateGroup(ctx context.Context, group account.Group) error
	DeleteGroup(ctx context.Context, groupID uuid.UUID) error
	Login(ctx context.Context, details LoginRequest) (*account.User, error)
	CreateContact(ctx context.Context, contact account.Contacts) (*account.Contacts, error)
	BlockContact(ctx context.Context, userID, contactID uuid.UUID) error
	UnblockContact(ctx context.Context, userID, contactID uuid.UUID) error
	GetUserContacts(ctx context.Context, contactID uuid.UUID) ([]account.Contacts, error)
	GetBlockedContacts(ctx context.Context, userID uuid.UUID) ([]account.Contacts, error)
	RemoveFavouriteContact(ctx context.Context, contactID, userID uuid.UUID) error
	MakeContactFavourite(ctx context.Context, contactID, userID uuid.UUID) error
	DeleteContact(ctx context.Context, contactID, userID uuid.UUID) error
	Ping(ctx context.Context, message string) string
	CreateChannel(ctx context.Context, channel account.Channel) (*account.Channel, error)
	DeleteChannel(ctx context.Context, channelID uuid.UUID) error
	AddUser(ctx context.Context, groupID, userID uuid.UUID) error
	RemoveUserFromGroup(ctx context.Context, groupID, userID uuid.UUID) error
	AddUserToChannel(ctx context.Context, userID, channelID uuid.UUID) error
}

type App struct {
	env               *env.Environment
	red               *connection.RedisOps
	kafka             *connection.Kafka
	logger            zerolog.Logger
	dbConnection      *gorm.DB
	userStore         store.UserDatabase
	groupStore        store.GroupDatabase
	channelStore      store.ChannelDatabase
	contactStore      store.ContactDatabase
	groupUserStore    store.GroupUserDatabase
	channelUserStore  store.ChannelUserDatabase
	settingsStore     store.SettingsDatabase
	userNoteStore     store.UserNoteDatabase
	userPasswordStore store.UserPasswordDatabase
	lastSeen          store.LastSeenDatabase
}

func New(z *zerolog.Logger, e *env.Environment) Operations {
	logger := z.With().Str(constants.PackageStrHelper, packageName).Logger()

	//red := connection.NewRedis(z, e)
	//db := connection.NewDBConn(*z, e)
	//kafka := connection.NewKafka(*z, e)
	//
	//userStore := store.NewUser(db)
	//groupStore := store.NewGroup(db)
	//channelStore := store.NewChannel(db)
	//contactStore := store.NewContact(db)
	//groupUserStore := store.NewGroupUser(db)
	//channelUserStore := store.NewChannelUser(db)
	//settingsStore := store.NewSettings(db)
	//userNoteStore := store.NewUserNote(db)
	//userPasswordStore := store.NewUserPassword(db)
	//lastSeenStore := store.NewLastSeen(db)

	app := &App{
		//red:               red,
		env:    e,
		logger: logger,
		//userStore:         *userStore,
		//groupStore:        *groupStore,
		//channelStore:      *channelStore,
		//contactStore:      *contactStore,
		//groupUserStore:    *groupUserStore,
		//channelUserStore:  *channelUserStore,
		//settingsStore:     *settingsStore,
		//userNoteStore:     *userNoteStore,
		//userPasswordStore: *userPasswordStore,
		//lastSeen:          *lastSeenStore,
		//dbConnection:      db.Connection,
		//kafka:             kafka,
	}

	logger.Info().Msg("application successfully initialized")

	return Operations(app)
}
