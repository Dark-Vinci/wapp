package chats

import (
	"time"

	"github.com/google/uuid"
)

type UserChatSettings struct {
	ID              uuid.UUID
	ReadReceipt     bool
	ReceivedReceipt bool
	UserID          uuid.UUID
	CreatedAt       time.Time
	DeletedAt       *time.Time
}
