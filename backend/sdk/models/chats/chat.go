package chats

import (
	"time"

	"github.com/google/uuid"
)

type UserMessage struct {
	ID            uuid.UUID
	FromUser      uuid.UUID
	ToUser        uuid.UUID
	FromMessageID *uuid.UUID
	ForwardedFrom *uuid.UUID
	ForwardedTo   *uuid.UUID
	ForwardedBy   *uuid.UUID
	Content       string
	Received      *time.Time
	Viewed        *time.Time
	Starred       *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

type GroupMessage struct {
	ID             uuid.UUID
	GroupID        uuid.UUID
	FromUser       uuid.UUID
	FromMessageID  *uuid.UUID
	Content        string
	Starred        bool
	PictureGroupID *uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

type ChannelMessage struct {
	ID             uuid.UUID
	ChannelID      uuid.UUID
	FromUser       uuid.UUID
	FromMessageID  *uuid.UUID // replying to a chat
	Content        string
	Starred        bool
	PictureGroupID *uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}
