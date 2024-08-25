package chats

import (
	"github.com/google/uuid"
	"time"
)

type Call struct {
	ID        uuid.UUID
	CallerID  uuid.UUID
	CreatedAt time.Time
	EndedAt   time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type UserCall struct {
	UserID    uuid.UUID
	CallID    uuid.UUID
	AddedBy   *uuid.UUID
	CreatedBy *uuid.UUID
}
