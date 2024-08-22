package chats

import (
	"time"

	"github.com/google/uuid"
)

type UserChat struct {
	ID         uuid.UUID
	FromUser   uuid.UUID
	ToUser     uuid.UUID
	FromChatID *uuid.UUID
	Content    string
	Received   bool
	Viewed     bool
	Starred    bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

type GroupChat struct {
	ID             uuid.UUID
	GroupID        uuid.UUID
	FromUser       uuid.UUID
	FromChatID     *uuid.UUID
	Content        string
	Starred        bool
	PictureGroupID *uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

type ChannelChat struct {
	ID             uuid.UUID
	ChannelID      uuid.UUID
	FromUser       uuid.UUID
	FromChatID     *uuid.UUID // replying to a chat
	Content        string
	Starred        bool
	PictureGroupID *uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}
