package server

import (
	"github.com/rs/zerolog"

	"github.com/dark-vinci/linkedout/backend/account/app"
	"github.com/dark-vinci/linkedout/backend/sdk/grpc/account"
	"github.com/dark-vinci/linkedout/backend/sdk/models"
)

const packageName = "server"

// Server struct
type Server struct {
	account.UnimplementedAccountServer
	env    *models.Env
	logger zerolog.Logger
	app    app.Operations
}

// New creates a new instance of the Server struct
func New(e *models.Env, z zerolog.Logger, a app.Operations) *Server {
	log := z.With().Str("PACKAGE_NAME", packageName).Logger()

	return &Server{
		env:    e,
		logger: log,
		app:    a,
	}
}
