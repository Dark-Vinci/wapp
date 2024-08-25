package server

import (
	"github.com/rs/zerolog"

	"github.com/dark-vinci/wapp/backend/post/app"
	"github.com/dark-vinci/wapp/backend/post/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/grpc/post"
)

const packageName = "post.server"

// Server struct
type Server struct {
	post.UnimplementedPostServer
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
