package store

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/post/connections"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

type SelectedContacts struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source selectedcontacts.go -destination ./mock/selectedcontacts_mock.go -package mock SelectedContactsDatabase
type SelectedContactsDatabase interface{}

func NewSelectedContacts(conn *connections.DBConn) *SelectedContactsDatabase {
	logger := conn.Log.With().Str(constants.PackageStrHelper, packageName).Logger()

	selectedContacts := &SelectedContacts{
		logger: &logger,
		db:     conn.Connection,
	}

	selectedContactsOperations := SelectedContactsDatabase(selectedContacts)

	logger.Info().Msg("successfully initialized selectedContacts database")

	return &selectedContactsOperations
}
