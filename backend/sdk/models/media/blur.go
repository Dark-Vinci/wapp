package media

import (
	"time"

	"github.com/google/uuid"
)

type Blur struct {
	ID        uuid.UUID
	URL       string
	MediaID   string
	For       Media
	CreatedAt time.Time
}
