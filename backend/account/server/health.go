package server

import (
	"context"

	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/grpc/account"
)

// Ping for service health checks
func (s *Server) Ping(ctx context.Context, in *account.PingRequest) (*account.PingResponse, error) {
	s.logger.Info().
		Str(constants.MethodStrHelper, "server.Ping").
		Str(constants.FunctionInputHelper, in.Data).
		Msg("got ping account service request")

	return &account.PingResponse{
		Data: s.app.Ping(ctx, in.GetData()),
	}, nil
}
