package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"primary key"`
	FirstName  string
	LastName   string
	MiddleName *string

	Email       string
	PhoneNumber string
	Password    *string

	About string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
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
