package server

import (
	"context"
	"fmt"

	"github.com/dark-vinci/linkedout/backend/sdk/constants"
	"github.com/dark-vinci/linkedout/backend/sdk/grpc/account"
)

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
