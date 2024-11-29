package account

import (
	"errors"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/dark-vinci/wapp/backend/gateway/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/grpc/account"
)

const packageName = "gateway.downstream.account"

type Account struct {
	client      account.AccountClient
	connection  *grpc.ClientConn
	env         *env.Environment
	logger      *zerolog.Logger
	isConnected bool
}

func New(e *env.Environment) *Account {
	return &Account{
		env: e,
	}
}

func (p *Account) Connect() error {
	log := p.logger.With().Str(constants.MethodStrHelper, "account.Connect").Logger()

	conn, err := grpc.Dial(p.env.AccountGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// try again in the next n minutes
		log.Err(err).Msg("account.Connect; fail to dial")
		return err
	}

	p.connection = conn
	p.isConnected = true

	return nil
}

func (p *Account) Client() (*Account, error) {
	if !p.isConnected {
		err := p.Connect()

		if err != nil {
			return nil, errors.New("post.Client; fail to connect")
		}
	}

	return p, nil
}

func (p *Account) Close() {
	_ = p.connection.Close()
	p.isConnected = false
}
