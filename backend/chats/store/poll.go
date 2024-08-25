package store

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/chats/connections"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

type Poll struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source poll.go -destination ./mock/poll_mock.go -package mock PollDatabase
type PollDatabase interface {
	//Create
}

func NewPollDatabase(conn *connections.DBConn) *PollDatabase {
	logger := conn.Logger.With().Str(constants.PackageStrHelper, packageName).Logger()

	poll := &Poll{
		logger: &logger,
		db:     conn.Connection,
	}

	pollOperations := PollDatabase(poll)

	logger.Info().Msg("successfully initialized poll database")

	return &pollOperations
}
