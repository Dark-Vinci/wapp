package media

import (
	"github.com/google/uuid"
	"time"
)

type Profile struct {
	Entity    string //TODO: make it enum of person, channel, group
	URL       string
	EntityID  uuid.UUID
	CreatedAt time.Time
	DeletedAt *time.Time
}
