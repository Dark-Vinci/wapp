package model

import "github.com/google/uuid"

type FileContent struct {
	UploadID string            `json:"upload_id"`
	Key      string            `json:"key"`
	Name     string            `json:"name"`
	Size     int64             `json:"size"`
	FileType string            `json:"type"`
	Headers  map[string]string `json:"headers"`
	UserID   uuid.UUID         `json:"user_id"`
}
