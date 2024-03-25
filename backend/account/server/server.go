package server

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"

	"github.com/dark-vinci/linkedout/backend/account/app"
	"github.com/dark-vinci/linkedout/backend/sdk/constants"
	"github.com/dark-vinci/linkedout/backend/sdk/grpc/account"
	"github.com/dark-vinci/linkedout/backend/sdk/models"
)

const packageName = "account.server"

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

// Ping for service health checks
func (s *Server) Ping(_ context.Context, in *account.PingRequest) (*account.PingResponse, error) {
	s.logger.Info().
		Str(constants.MethodStrHelper, "server.Ping").
		Str(constants.FunctionInputHelper, in.Data).
		Msg("got ping account service request")

	return &account.PingResponse{
		Data: fmt.Sprintf("%s says %s", packageName, in.GetData()),
	}, nil
}
