package media

import (
	"github.com/google/uuid"
	"time"
)

type PostMedia struct {
	UserID    uuid.UUID
	URL       string
	MediaType string
	CreatedAt time.Time
	DeletedAt *time.Time
}
