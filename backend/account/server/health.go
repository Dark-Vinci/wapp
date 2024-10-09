package server

import (
	"context"

	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/grpc/account"
)

// Ping for service health checks
func (s *Server) Ping(ctx context.Context, in *account.PingRequest) (*account.PingResponse, error) {
	_, span := s.tracer.Start(ctx, "account.Ping")
	defer span.End()

	s.logger.Info().
		Str(constants.MethodStrHelper, "server.Ping").
		Str(constants.FunctionInputHelper, in.Data).
		Str(constants.RequestID, in.GetData()).
		Msg("got ping account service request")

	ctx = context.WithValue(ctx, constants.RequestID, in.GetRequestID())

	return &account.PingResponse{
		Data: s.app.Ping(ctx, in.GetData()),
	}, nil
}
