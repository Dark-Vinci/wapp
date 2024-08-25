package store

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/chats/connections"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

type PollOption struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source polloption.go -destination ./mock/polloption_mock.go -package mock PollOptionDatabase
type PollOptionDatabase interface{}

func NewPollOptionDatabase(conn *connections.DBConn) *PollOptionDatabase {
	logger := conn.Logger.With().Str(constants.PackageStrHelper, packageName).Logger()

	pollOption := &PollOption{
		logger: &logger,
		db:     conn.Connection,
	}

	pollOptionOperations := PollOptionDatabase(pollOption)

	logger.Info().Msg("successfully initialized pollVote database")

	return &pollOptionOperations
}
