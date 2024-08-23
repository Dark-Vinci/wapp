package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/account/connection"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models/account"
)

type Settings struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source settings.go -destination ./mock/settings_mock.go -package mock SettingsDatabase
type SettingsDatabase interface {
	Create(ctx context.Context, settings account.Settings) (*account.Settings, error)
	GetUserSettings(ctx context.Context, userID uuid.UUID) (*account.Settings, error)
	Update(ctx context.Context, settings account.Settings) (*account.Settings, error)
}

func NewSettings(conn *connection.DBConn) *SettingsDatabase {
	logger := conn.Log.With().
		Str(constants.FunctionNameHelper, "NewSettings").
		Str(constants.PackageStrHelper, packageName).
		Logger()

	settings := &Settings{
		logger: &logger,
		db:     conn.Connection,
	}

	settingsOperations := SettingsDatabase(settings)

	return &settingsOperations
}

func (s *Settings) Create(ctx context.Context, settings account.Settings) (*account.Settings, error) {
	return nil, nil
}

func (s *Settings) GetUserSettings(ctx context.Context, userID uuid.UUID) (*account.Settings, error) {
	return nil, nil
}

func (s *Settings) Update(ctx context.Context, settings account.Settings) (*account.Settings, error) {
	return nil, nil
}
