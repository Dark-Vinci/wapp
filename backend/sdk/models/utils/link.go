package utils

import (
	"time"

	"github.com/google/uuid"
)

type Link struct {
	ID        uuid.UUID
	Initial   string
	Generated string
	CreatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
