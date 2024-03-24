package account

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	ID        uuid.UUID
	Name      string
	CreatedBy time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
