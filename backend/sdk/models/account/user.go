package account

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"primary key"`
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
	LockChat  bool
	//ProfileURL *string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Channel struct {
	ID        uuid.UUID `gorm:"primary key"`
	UserID    uuid.UUID `gorm:"index"`
	Name      string
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

type Contacts struct {
	ID            string
	Owner         string
	ContactID     string
	ContactNumber string
	IsVerified    bool
	IsBlocked     bool
	IsArchived    bool
	IsFavorite    bool
	createdAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}
