package chats

import (
	"time"

	"github.com/google/uuid"
)

type ChatReaction struct {
	ChatID    uuid.UUID
	Count     uint
	CreatedAt time.Time
	DeletedAt *time.Time
}

type UserChatReaction struct {
	ChatReactionID uuid.UUID
	Emoji          string
	UserID         uuid.UUID
	CreatedAt      time.Time
	DeletedAt      *time.Time
}
