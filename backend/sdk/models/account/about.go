package account

import (
	"time"

	"github.com/google/uuid"
)

type About struct {
	ID          uuid.UUID
	EntityID    uuid.UUID
	Content     string
	EntityType  string
	CompanyURL  string
	PhoneNumber string
	Industry    string
	FoundedIn   time.Time
	Specialties string
	CompanySize string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
