package posts

import (
	"github.com/google/uuid"
	"time"
)

type ContactPost struct {
	UserID     uuid.UUID
	ConsumerID uuid.UUID
	Muted      bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
