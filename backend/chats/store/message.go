package store

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/chats/connections"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

type Message struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source message.go -destination ./mock/message_mock.go package mock MessageDatabase
type MessageDatabase interface{}

func NewMessageDatabase(conn *connections.DBConn) *MessageDatabase {
	logger := conn.Logger.With().
		Str(constants.FunctionNameHelper, "NewMessageDatabase").
		Str(constants.PackageStrHelper, packageName).
		Logger()

	chat := &Message{}

	chatOperations := MessageDatabase(chat)

	logger.Info().Msg("successfully initialized chat database")

	return &chatOperations
}

//func
