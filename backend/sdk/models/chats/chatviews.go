package chats

import (
	"time"

	"github.com/google/uuid"
)

type GroupChatView struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	GroupChatID uuid.UUID
	Viewer      uuid.UUID
	CreatedAt   time.Time
	For         string // channel, group
}

//type ChannelChatView struct {
//	ID          uuid.UUID
//	UserID      uuid.UUID
//	GroupChatID uuid.UUID
//	Viewer      uuid.UUID
//	CreatedAt   time.Time
//}
