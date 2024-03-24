package account

import (
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	UserContactID *uuid.UUID
	Value         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAT     *time.Time
}
