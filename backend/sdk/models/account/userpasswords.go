package account

import (
	"github.com/google/uuid"
	"time"
)

type UserPasswords struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	ClientName string
	Password   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

func (u *UserPasswords) From(payload CreatePasswordPayload) {
	u.ClientName = payload.ClientName
	u.Password = payload.Password
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.UserID = payload.UserID
}

type CreatePasswordPayload struct {
	UserID     uuid.UUID
	Password   string `json:"password"`
	ClientName string `json:"client_name"`
}

func a() {
	v := UserPasswords{}
	v.From(CreatePasswordPayload{})
}
