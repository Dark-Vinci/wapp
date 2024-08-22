package posts

import (
	"github.com/google/uuid"
	"time"
)

type PostView struct {
	PostID    uuid.UUID
	ViewerID  uuid.UUID
	CreatedAt time.Time
}
