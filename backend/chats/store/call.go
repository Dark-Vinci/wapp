package store

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/chats/connections"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

type Call struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source call.go -destination ./mock/call_mock.go -package mock CallDatabase
type CallDatabase interface{}

func NewCallDatabase(conn *connections.DBConn) *CallDatabase {
	logger := conn.Logger.With().Str(constants.PackageStrHelper, packageName).Logger()

	call := &Call{
		logger: &logger,
		db:     conn.Connection,
	}

	callOperations := CallDatabase(call)

	logger.Info().Msg("successfully initialized call database")

	return &callOperations
}
