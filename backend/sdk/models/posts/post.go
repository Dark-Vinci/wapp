package posts

import (
	"github.com/google/uuid"
	"time"
)

type MediaType int

const (
	// TEXT text, emoji, URL
	TEXT = iota + 1
	Audio
	Video
)

func (mt MediaType) String() string {
	switch mt {
	case Audio:
		return "audio"
	case Video:
		return "video"
	case TEXT:
		return "text"
	default:
		return "unknown"
	}
}

type Post struct {
	ID        uuid.UUID `gorm:"primary_key"`
	UserID    uuid.UUID `gorm:"index"`
	MediaID   uuid.UUID
	Type      MediaType
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
