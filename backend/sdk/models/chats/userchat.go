package chats

import (
	"github.com/google/uuid"
	"time"
)

type UserChat struct {
	ID        uuid.UUID
	ForUser   uuid.UUID
	WithUser  uuid.UUID
	CreatedAt time.Time
	UpdatedAt *time.Time
}
