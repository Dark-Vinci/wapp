package account

import (
	"github.com/google/uuid"
	"time"
)

type Group struct {
	ID        uuid.UUID
	Name      string
	About     string
	CreatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
