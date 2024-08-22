package chats

import (
	"github.com/google/uuid"
	"time"
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
