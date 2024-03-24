package account

import (
	"time"

	"github.com/google/uuid"
)

type Connections struct {
	ID          uuid.UUID
	InitiatedBy uuid.UUID
	ToUser      uuid.UUID
	Accepted    bool
	Discarded   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
