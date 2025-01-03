package media

import (
	"time"

	"github.com/google/uuid"
)

type PostMedia struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	PostID    uuid.UUID
	URL       string
	Type      Type // AUDIO, VIDEO
	CreatedAt time.Time
	DeletedAt *time.Time
}
