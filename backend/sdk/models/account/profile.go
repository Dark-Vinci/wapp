package account

import (
	"github.com/google/uuid"
	"time"
)

type Profile struct {
	ID uuid.UUID

	//enum of company, user, orgs
	EntityType string

	EntityID uuid.UUID

	About string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
