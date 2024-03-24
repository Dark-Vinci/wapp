package reactions

import (
	"time"

	"github.com/google/uuid"
)

type Repost struct {
	ID        uuid.UUID
	PostID    uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
