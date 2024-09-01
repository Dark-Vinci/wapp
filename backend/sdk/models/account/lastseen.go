package account

import (
	"github.com/google/uuid"
	"time"
)

type LastSeen struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
}
