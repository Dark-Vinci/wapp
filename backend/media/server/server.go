package server

import (
	"github.com/rs/zerolog"

	"github.com/dark-vinci/wapp/backend/media/app"
	"github.com/dark-vinci/wapp/backend/media/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/grpc/media"
)

const packageName = "media.server"

// Server struct
type Server struct {
	media.UnimplementedMediaServer
	env    *env.Environment
	logger zerolog.Logger
	app    app.Operations
}

// New creates a new instance of the Server struct
func New(e *env.Environment, z zerolog.Logger, a app.Operations) *Server {
	log := z.With().Str(constants.PackageStrHelper, packageName).Logger()

	return &Server{
		env:    e,
		logger: log,
		app:    a,
	}
}
