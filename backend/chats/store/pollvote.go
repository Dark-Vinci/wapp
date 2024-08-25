package store

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/chats/connections"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

type PollVote struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source pollvote.go -destination ./mock/pollvote_mock.go -package mock PollVoteDatabase
type PollVoteDatabase interface{}

func NewPollVoteDatabase(conn *connections.DBConn) *PollVoteDatabase {
	logger := conn.Logger.With().Str(constants.PackageStrHelper, packageName).Logger()

	pollVote := &PollVote{
		logger: &logger,
		db:     conn.Connection,
	}

	pollVoteOperations := PollVoteDatabase(pollVote)

	logger.Info().Msg("successfully initialized pollVote database")

	return &pollVoteOperations
}
