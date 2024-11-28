package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"

	"github.com/dark-vinci/wapp/backend/sdk/utils/crypto"
)

type User struct {
	ID         uuid.UUID `gorm:"primary key"`
	FirstName  string
	LastName   string
	MiddleName *string

	Email       string
	PhoneNumber string
	Password    string

	About string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	c := crypto.Crypto{Cost: 10}
	hash, err := c.HashPassword(u.Password)

	u.Password = hash
	return
}

type Group struct {
	ID   uuid.UUID `gorm:"primary key"`
	Name string

	CreatedBy uuid.UUID
	Pofile    *string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
