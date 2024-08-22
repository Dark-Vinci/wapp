package account

import (
	"github.com/google/uuid"
	"time"
)

type Channel struct {
	ID        uuid.UUID `gorm:"primary key"`
	UserID    uuid.UUID `gorm:"index"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
