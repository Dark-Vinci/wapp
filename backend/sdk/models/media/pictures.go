package media

import (
	"github.com/google/uuid"
	"time"
)

type PictureGroup struct {
	ID     uuid.UUID
	PostID uuid.UUID `json:"postId"` // chat id
	Count  uint8
}

type Pictures struct {
	PictureGroupID uuid.UUID
	URL            string `json:"url"`
	PostedBy       uuid.UUID
	EntityID       uuid.UUID
	Entity         string // chat, group, channel
	CreatedAt      time.Time
	DeletedAt      *time.Time
}
