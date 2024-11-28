package app

import (
	"context"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/wapp/backend/gateway/model"
	"github.com/dark-vinci/wapp/backend/sdk/utils/s3"
)

const packageName = "gateway.app"

type Operations interface {
	Ping() string
	CreateUser() string

	LoginToAccount(ctx context.Context, req model.LoginRequest) (string, error)
	UploadSingleMedia(ctx context.Context, userID uuid.UUID, file *multipart.FileHeader) (*model.FileContent, error)
	UploadMultipleMedia(ctx context.Context, files []*multipart.FileHeader) ([]model.FileContent, error)
}

type App struct {
	ss3    s3.MediaStore
	logger *zerolog.Logger
}

func New() Operations {
	app := &App{}

	return Operations(app)
}

func (a *App) Ping() string {
	return "healthy"
}

func (a *App) CreateUser() string {
	return "user"
}
