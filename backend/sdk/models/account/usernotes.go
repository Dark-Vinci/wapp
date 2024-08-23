package account

import (
	"github.com/google/uuid"
	"time"
)

type UserNotes struct {
	ID        uuid.UUID
	DeletedAt *time.Time
}
