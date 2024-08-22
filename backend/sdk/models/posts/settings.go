package posts

import (
	"github.com/google/uuid"
	"time"
)

type PostSettings struct {
	UserID     uuid.UUID
	Publicity  string // EVERYONE CONTACTS, SELECTED, EXCEPT
	CreatedAt  time.Time
	DeletedAt  *time.Time
	Visibility string //MAKE VIEWING COUNT, DONT MAKE VIEWING COUNT
}

type SelectedContacts struct {
	UserID         uuid.UUID // who should see it
	PostSettingsID uuid.UUID
	PosterID       uuid.UUID
	CreatedAt      time.Time
	DeletedAt      *time.Time
}

type ExceptContacts struct {
	UserID         uuid.UUID
	PostSettingsID uuid.UUID
	PosterID       uuid.UUID
	CreatedAt      time.Time
	DeletedAt      *time.Time
}
