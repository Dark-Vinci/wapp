package server

import (
	"context"

	"github.com/dark-vinci/wapp/backend/account/server/from"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/grpc/account"
)

func (s *Server) CreateUserAccount(ctx context.Context, in *account.CreateUserRequest) (*account.CreateUserResponse, error) {
	log := s.logger.With().
		Str(constants.MethodStrHelper, "server.CreateUserAccount").
		Str(constants.RequestID, in.GetRequestId()).Logger()

	ctx = context.WithValue(ctx, constants.RequestID, in.GetRequestId())

	userRequest := from.CreateUserRequest(in)

	user, err := s.app.CreateUser(ctx, userRequest)
	if err != nil {
		log.Err(err).Msg("Server: failed to create user")
		return nil, err
	}

	return from.CreateUserResponse(in.GetRequestId(), *user), nil
}

func (s *Server) LoginToAccount(ctx context.Context, in *account.LoginRequest) (*account.LoginResponse, error) {
	log := s.logger.With().
		Str(constants.MethodStrHelper, "server.LoginToAccount").
		Str(constants.RequestID, ctx.Value(constants.RequestID).(string)).Logger()

	ctx = context.WithValue(ctx, constants.RequestID, in.GetRequestId())

	if err := s.app.LoginToAccount(ctx, in.PhoneNumber, in.Password); err != nil {
		log.Err(err).Msg("Server: failed to login")
		return nil, err
	}

	return &account.LoginResponse{
		RequestId: in.GetRequestId(),
		Success:   true,
	}, nil
}
