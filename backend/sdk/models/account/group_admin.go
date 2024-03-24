package account

import (
	"time"

	"github.com/google/uuid"
)

type GroupAdmin struct {
	ID           uuid.UUID
	GroupID      uuid.UUID
	IsSuperAdmin bool
	MadeAdminBy  *uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
