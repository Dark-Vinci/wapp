package chats

import (
	"errors"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/dark-vinci/wapp/backend/gateway/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/grpc/chat"
)

const packageName = "gateway.downstream.chat"

type Chat struct {
	client      chat.ChatClient
	connection  *grpc.ClientConn
	env         *env.Environment
	logger      *zerolog.Logger
	isConnected bool
}

func New(e *env.Environment) *Chat {
	return &Chat{
		env: e,
	}
}

func (p *Chat) Connect() error {
	log := p.logger.With().Str(constants.MethodStrHelper, "chat.Connect").Logger()

	conn, err := grpc.Dial(p.env.ChatGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// try again in the next n minutes
		log.Err(err).Msg("chat.Connect; fail to dial")
		return err
	}

	p.connection = conn
	p.isConnected = true

	return nil
}
func (p *Chat) Client() (*Chat, error) {
	if !p.isConnected {
		err := p.Connect()

		if err != nil {
			return nil, errors.New("chat.Client; fail to connect")
		}
	}

	return p, nil
}

func (p *Chat) Close() {
	_ = p.connection.Close()
	p.isConnected = false
}
