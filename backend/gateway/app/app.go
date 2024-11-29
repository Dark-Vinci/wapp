package app

import (
	"context"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/wapp/backend/gateway/downstream"
	"github.com/dark-vinci/wapp/backend/gateway/env"
	"github.com/dark-vinci/wapp/backend/gateway/model"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/utils/mail"
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
	ss3        s3.MediaStore
	logger     *zerolog.Logger
	env        *env.Environment
	mailer     mail.Mailer
	downstream *downstream.Downstream
}

func New(z *zerolog.Logger, e *env.Environment) Operations {
	dst := downstream.New(z, e)
	ss3 := s3.NewS3("my-region") // todo; update accordingly
	mailer := mail.New()

	logger := z.With().Str(constants.PackageStrHelper, packageName).Logger()

	app := &App{
		downstream: dst,
		logger:     &logger,
		env:        e,
		ss3:        ss3,
		mailer:     mailer,
	}

	return Operations(app)
}

func (a *App) Ping() string {
	return "healthy"
}

func (a *App) CreateUser() string {
	return "user"
}
