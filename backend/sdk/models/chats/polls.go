package chats

import (
	"github.com/google/uuid"
	"time"
)

type Poll struct {
	ID           uuid.UUID `json:"id"`
	PostedBy     uuid.UUID
	Title        string
	SingleChoice bool
	ChatID       uuid.UUID
	VoteCount    uint32
	Entity       string // chat, group, channel
}

type PollsOptions struct {
	ID     uuid.UUID
	PollID uuid.UUID
	Choice string
}

type PollVote struct {
	UserID    uuid.UUID
	PollID    uuid.UUID
	OptionID  uuid.UUID
	CreatedAt time.Time
	DeletedAt *time.Time
}
