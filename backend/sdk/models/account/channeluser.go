package account

import (
	"github.com/google/uuid"
	"time"
)

type ChannelUser struct {
	ID        uuid.UUID `gorm:"primary key"`
	ChannelID uuid.UUID `gorm:"index"`
	UserID    uuid.UUID `gorm:"index"`
	Mute      bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
