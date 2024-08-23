package chats

import (
	"time"

	"github.com/google/uuid"
)

type Forward struct {
	FromChatID  uuid.UUID // Group, channel, chat
	ForwardedBy uuid.UUID
	ToChatID    uuid.UUID // Group, channel, chat
	ChatID      uuid.UUID // the content
	CreatedAt   time.Time
	DeletedAt   *time.Time
}
