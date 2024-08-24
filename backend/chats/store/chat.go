package store

import (
	"github.com/dark-vinci/wapp/backend/chats/connections"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Chat struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

type ChatDatabase interface{}

func NewChatDatabase(conn *connections.DBConn) *ChatDatabase {
	logger := conn.Logger.With().
		Str(constants.FunctionNameHelper, "NewChatDatabase").
		Str(constants.PackageStrHelper, packageName).
		Logger()

	chat := &Chat{}

	chatOperations := ChatDatabase(chat)

	logger.Info().Msg("successfully initialized chat database")

	return &chatOperations
}
