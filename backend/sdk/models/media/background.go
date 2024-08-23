package media

import (
	"github.com/google/uuid"
	"time"
)

type Background struct {
	UserID    uuid.UUID
	URL       string
	CreatedAt time.Time
}
