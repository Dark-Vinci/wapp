package account

import (
	"github.com/google/uuid"
	"time"
)

type GroupUser struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	GroupID   uuid.UUID
	Mute      bool
	IsOwner   bool
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
