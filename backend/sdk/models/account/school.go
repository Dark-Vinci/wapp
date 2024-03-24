package account

import (
	"github.com/google/uuid"
	"time"
)

type School struct {
	ID    uuid.UUID
	Name  string
	About string
	//CreatedAt string
	FoundedBy string
	FoundedAt time.Time
	CreatedAt string
	UpdatedAt string
	DeletedAt *time.Time
	Email     string
	Password  string
}
