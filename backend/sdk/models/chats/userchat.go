package chats

import (
	"time"

	"github.com/google/uuid"
)

type UserChatj struct {
	ID        uuid.UUID
	ForUser   uuid.UUID
	WithUser  uuid.UUID
	CreatedAt time.Time
	UpdatedAt *time.Time
}
