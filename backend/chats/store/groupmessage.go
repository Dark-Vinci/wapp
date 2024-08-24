package store

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/chats/connections"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

type GroupMessage struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source groupmessage.go -destination ./mock/groupmessage_mock.go -package mock GroupMessageDatabase
type GroupMessageDatabase interface{}

func NewGroupMessageDatabase(conn *connections.DBConn) *GroupMessageDatabase {
	logger := conn.Logger.With().
		Str(constants.FunctionNameHelper, "NewGroupMessage").
		Str(constants.PackageStrHelper, packageName).
		Logger()

	groupMessage := &GroupMessage{
		logger: &logger,
		db:     conn.Connection,
	}

	groupMessageOperations := GroupMessageDatabase(groupMessage)

	return &groupMessageOperations
}
