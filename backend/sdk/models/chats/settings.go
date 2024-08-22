package chats

import "github.com/google/uuid"

type UserChatSettings struct {
	ReadReceipt     bool
	ReceivedReceipt bool
	UserID          uuid.UUID
	ID              uuid.UUID
}
