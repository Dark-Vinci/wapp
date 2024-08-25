package store

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/post/connections"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

type Settings struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source settings.go -destination ./mock/settings_mock.go -package mock SettingsDatabase
type SettingsDatabase interface{}

func NewSettings(conn *connections.DBConn) *SettingsDatabase {
	logger := conn.Log.With().Str(constants.PackageStrHelper, packageName).Logger()

	post := &Settings{
		logger: &logger,
		db:     conn.Connection,
	}

	postOperations := SettingsDatabase(post)

	logger.Info().Msg("successfully initialized settings database")

	return &postOperations
}
