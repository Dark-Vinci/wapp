package store

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/chats/connections"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

type Settings struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source settings.go -destination ./mock/settings_mock.go -package mock SettingsDatabase
type SettingsDatabase interface{}

func NewSettingsDatabase(conn *connections.DBConn) *SettingsDatabase {
	logger := conn.Logger.With().Str(constants.PackageStrHelper, packageName).Logger()

	settings := &Settings{
		logger: &logger,
		db:     conn.Connection,
	}

	settingsOperations := SettingsDatabase(settings)

	logger.Info().Msg("successfully initialized settings database")

	return &settingsOperations
}
