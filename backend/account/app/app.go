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

type Operations interface {
	DeleteUserAccount(ctx context.Context, userID uuid.UUID) error
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	CreateGroup(ctx context.Context, group account.Group) error
	DeleteGroup(ctx context.Context, groupID uuid.UUID) error
	Login(ctx context.Context, userID uuid.UUID) (*models.User, error)
	CreateContact(ctx context.Context, contact account.Contacts) (*account.Contacts, error)
	BlockContact(ctx context.Context, userID, contactID uuid.UUID) error
	UnblockContact(ctx context.Context, userID, contactID uuid.UUID) error
	GetUserContacts(ctx context.Context, contactID uuid.UUID) ([]account.Contacts, error)
	GetBlockedContacts(ctx context.Context, userID uuid.UUID) ([]account.Contacts, error)
	RemoveFavouriteContact(ctx context.Context, contactID, userID uuid.UUID) error
	MakeContactFavourite(ctx context.Context, contactID, userID uuid.UUID) error
	DeleteContact(ctx context.Context, contactID, userID uuid.UUID) error
	Ping(ctx context.Context, message string) string
}

type App struct {
	env          *env.Environment
	red          *connection.RedisClient
	kafka        *connection.Kafka
	logger       zerolog.Logger
	dbConnection *gorm.DB
	userStore    store.UserDatabase
	groupStore   store.GroupDatabase
	channelStore store.ChannelDatabase
	contactStore store.ContactDatabase
}

func New(z *zerolog.Logger, e *env.Environment) Operations {
	logger := z.With().Str(constants.PackageStrHelper, packageName).Logger()

	red := connection.NewRedisClient(z, e)
	db := connection.NewDBConn(*z, e)
	kafka := connection.NewKafka(*z, e)

	userStore := store.NewUser(db)
	groupStore := store.NewGroup(db)
	channelStore := store.NewChannel(db)
	contactStore := store.NewContact(db)

	app := &App{
		red:          red,
		env:          e,
		logger:       logger,
		userStore:    *userStore,
		groupStore:   *groupStore,
		channelStore: *channelStore,
		contactStore: *contactStore,
		dbConnection: db.Connection,
		kafka:        kafka,
	}

	return Operations(app)
}
