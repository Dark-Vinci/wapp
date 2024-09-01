package server

import (
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"

	"github.com/dark-vinci/wapp/backend/account/app"
	"github.com/dark-vinci/wapp/backend/account/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/grpc/account"
)

const packageName = "account.server"

// Server struct
type Server struct {
	account.UnimplementedAccountServer
	env    *env.Environment
	logger zerolog.Logger
	app    app.Operations
	tracer trace.Tracer
}

// New creates a new instance of the Server struct
func New(e *env.Environment, z zerolog.Logger, a app.Operations, t trace.Tracer) *Server {
	log := z.With().Str(constants.PackageStrHelper, packageName).Logger()

	return &Server{
		env:    e,
		logger: log,
		app:    a,
		tracer: t,
	}
}
