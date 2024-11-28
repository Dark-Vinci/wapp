package media

import (
	"github.com/google/uuid"
	"time"
)

type Profile struct {
	ID  uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	URL string
	//UserID    uuid.UUID
	AccountID uuid.UUID // Group, User, Channel
	CreatedBy uuid.UUID // Group{user can be different}
	CreatedAt time.Time
	DeletedAt *time.Time
}

//func (p *Profile) from(userID uuid.UUID, URL string) string {
//
//}
