package media

import (
	"errors"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/dark-vinci/wapp/backend/gateway/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/grpc/media"
)

const packageName = "gateway.downstream.media"

type Media struct {
	client      media.MediaClient
	connection  *grpc.ClientConn
	env         *env.Environment
	logger      *zerolog.Logger
	isConnected bool
}

func New(e *env.Environment) *Media {
	return &Media{
		env: e,
	}
}

func (p *Media) Connect() error {
	if p.isConnected {
		return nil
	}

	log := p.logger.With().Str(constants.MethodStrHelper, "media.Connect").Logger()

	conn, err := grpc.Dial(p.env.MediaGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// try again in the next n minutes
		log.Err(err).Msg("media.Connect; fail to dial")
		return err
	}

	p.connection = conn
	p.isConnected = true

	return nil
}

func (p *Media) Client() (*Media, error) {
	if !p.isConnected {
		err := p.Connect()

		if err != nil {
			return nil, errors.New("media.Client; fail to connect")
		}
	}

	return p, nil
}

func (p *Media) Close() {
	_ = p.connection.Close()
	p.isConnected = false
}
