package model

import (
	"time"

	"github.com/google/uuid"
)

type WSMessage struct {
	Data      string   `json:"data"`
	Media     []string `json:"media"`
	userID    uuid.UUID
	to        uuid.UUID
	toType    string    // enum: personal, group, channel, post
	Timestamp time.Time `json:"timestamp"`
}
