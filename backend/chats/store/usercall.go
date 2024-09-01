package store

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/chats/connections"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

type UserCall struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source usercall.go -destination ./mock/usercall_mock.go -package mock UserCallDatabase
type UserCallDatabase interface{}

func NewUserCallDatabase(conn *connections.DBConn) *UserCallDatabase {
	logger := conn.Logger.With().Str(constants.PackageStrHelper, packageName).Logger()

	userCall := &UserCall{
		logger: &logger,
		db:     conn.Connection,
	}

	userCallOperations := UserCallDatabase(userCall)

	logger.Info().Msg("successfully initialized pollVote database")

	return &userCallOperations
}
