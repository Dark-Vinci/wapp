package post

import (
	"errors"
	"github.com/dark-vinci/wapp/backend/gateway/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/grpc/post"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const packageName = "gateway.downstream.post"

type Post struct {
	client      post.PostClient
	connection  *grpc.ClientConn
	env         *env.Environment
	logger      *zerolog.Logger
	isConnected bool
}

func New(e *env.Environment) *Post {
	return &Post{
		env: e,
	}
}

func (p *Post) Connect() error {
	log := p.logger.With().Str(constants.MethodStrHelper, "post.Connect").Logger()

	conn, err := grpc.Dial(p.env.PostGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// try again in the next n minutes
		log.Err(err).Msg("post.Connect; fail to dial")
		return err
	}

	p.connection = conn
	p.isConnected = true

	return nil
}

func (p *Post) Client() (*Post, error) {
	if !p.isConnected {
		err := p.Connect()

		if err != nil {
			return nil, errors.New("post.Client; fail to connect")
		}
	}

	return p, nil
}

func (p *Post) Close() {
	_ = p.connection.Close()
	p.isConnected = false
}
