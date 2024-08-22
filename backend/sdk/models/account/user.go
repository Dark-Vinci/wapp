package account

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"primary key"`
	Username    *string
	FirstName   string
	LastName    string
	MiddleName  *string
	Email       string
	PhoneNumber string
	Password    *string
	About       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type CreateUser struct {
	FirstName   string
	LastName    string
	MiddleName  *string
	Email       string
	PhoneNumber string
	Password    *string
	About       string
}

type ProfileURL struct {
	ID        uuid.UUID
	EntityID  uuid.UUID
	URL       string
	FOR       string // channel, user, group
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Group struct {
	ID        uuid.UUID `gorm:"primary key"`
	Name      string
	CreatedBy uuid.UUID
	LockChat  bool //only an admin can post
	//ProfileURL *string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type UserGroup struct {
	ID        uuid.UUID `gorm:"primary key"`
	GroupID   uuid.UUID `gorm:"index"`
	UserID    uuid.UUID `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Mute      bool
	Deleted   bool
}
