package store

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/post/connections"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

type ContactPost struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source contactpost.go -destination ./mock/contactpost_mock.go -package mock ContactPostDatabase
type ContactPostDatabase interface{}

func NewContactPost(conn *connections.DBConn) *ContactPostDatabase {
	logger := conn.Log.With().Str(constants.PackageStrHelper, packageName).Logger()

	contactPost := &ContactPost{
		logger: &logger,
		db:     conn.Connection,
	}

	contactPostOperations := ContactPostDatabase(contactPost)

	logger.Info().Msg("successfully initialized contactPost database")

	return &contactPostOperations
}
