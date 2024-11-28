package media

import (
	"time"

	"github.com/google/uuid"
)

type Type string

var (
	Image Type = "image"
	Video Type = "video"
	Audio Type = "audio"
	GIF   Type = "gif"
	DOCS  Type = "docs" // pdf, excel, and others
)

type ChatMedia struct {
	ID        uuid.UUID
	PostedBy  uuid.UUID
	To        uuid.UUID
	Type      Type
	For       Media // GROUP, CHAT, CHANNEL
	MessageID uuid.UUID
	URL       uuid.UUID
	Once      bool
	CreatedAt time.Time
	DeletedAt *time.Time
}

type MultiUserChatMedia struct {
	For       string // group, channel
	ID        uuid.UUID
	MessageID uuid.UUID
	Once      bool
	URL       string
	Users     []SpecifiedUserMultiUserChatMedia `gorm:"foreignKey:MultiUserChatMediaID;references:MultiUserChatMediaID"`
}

type SpecifiedUserMultiUserChatMedia struct {
	ID                   uuid.UUID
	ForUserID            uuid.UUID
	MultiUserChatMediaID uuid.UUID
	MessageID            uuid.UUID
	Once                 bool
	Seen                 bool
}

type UserChatMedia struct {
	ID          uuid.UUID `json:"id"`
	ToUser      uuid.UUID
	Seen        bool
	ChatMediaID uuid.UUID
}
