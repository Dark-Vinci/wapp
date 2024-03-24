package reactions

import (
	"time"

	"github.com/google/uuid"
)

type Comments struct {
	ID        uuid.UUID
	PostID    uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
