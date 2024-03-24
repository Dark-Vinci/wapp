package utils

import (
	"time"

	"github.com/google/uuid"
)

type Search struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
