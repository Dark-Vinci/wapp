package account

import "time"

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
	DeletedAt     *time.Time
}
