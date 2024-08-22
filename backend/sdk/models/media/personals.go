package media

import (
	"github.com/google/uuid"
	"time"
)

type Personals struct {
	URL       string `json:"url"`
	UserID    uuid.UUID
	CreatedAt time.Time
	DeletedAt *time.Time
}
