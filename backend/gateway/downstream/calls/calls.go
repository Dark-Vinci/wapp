package calls

import (
	"errors"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/dark-vinci/wapp/backend/gateway/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/grpc/calls"
)

const packageName = "gateway.downstream.calls"

type Call struct {
	client      calls.CallClient
	connection  *grpc.ClientConn
	env         *env.Environment
	logger      *zerolog.Logger
	isConnected bool
}

func New(e *env.Environment) *Call {
	return &Call{
		env: e,
	}
}

func (p *Call) Connect() error {
	log := p.logger.With().Str(constants.MethodStrHelper, "calls.Connect").Logger()

	conn, err := grpc.Dial(p.env.CallGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// try again in the next n minutes
		log.Err(err).Msg("calls.Connect; fail to dial")
		return err
	}

	p.connection = conn
	p.isConnected = true

	return nil
}

func (p *Call) Client() (*Call, error) {
	if !p.isConnected {
		err := p.Connect()

		if err != nil {
			return nil, errors.New("post.Client; fail to connect")
		}
	}

	return p, nil
}

func (p *Call) Close() {
	_ = p.connection.Close()
	p.isConnected = false
}
