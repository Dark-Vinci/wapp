package server

import (
	"github.com/rs/zerolog"

	"github.com/dark-vinci/linkedout/backend/account/app"
	"github.com/dark-vinci/linkedout/backend/account/env"
	"github.com/dark-vinci/linkedout/backend/sdk/grpc/account"
)

const packageName = "account.server"
const packageNameKey string = "PACKAGE_NAME"

// Server struct
type Server struct {
	account.UnimplementedAccountServer
	env    *env.Environment
	logger zerolog.Logger
	app    app.Operations
}

// New creates a new instance of the Server struct
func New(e *env.Environment, z zerolog.Logger, a app.Operations) *Server {
	log := z.With().Str(packageNameKey, packageName).Logger()

	return &Server{
		env:    e,
		logger: log,
		app:    a,
	}
}
