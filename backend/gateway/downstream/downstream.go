package downstream

import (
	"github.com/rs/zerolog"

	"github.com/dark-vinci/wapp/backend/gateway/downstream/account"
	"github.com/dark-vinci/wapp/backend/gateway/downstream/calls"
	"github.com/dark-vinci/wapp/backend/gateway/downstream/chats"
	"github.com/dark-vinci/wapp/backend/gateway/downstream/media"
	"github.com/dark-vinci/wapp/backend/gateway/downstream/post"
	"github.com/dark-vinci/wapp/backend/gateway/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

const packageName = "gateway.downstream"

type Downstream struct {
	logger  *zerolog.Logger
	env     *env.Environment
	Account *account.Account
	Post    *post.Post
	Chat    *chats.Chat
	Call    *calls.Call
	Media   *media.Media
}

func New(z *zerolog.Logger, e *env.Environment) *Downstream {
	logger := z.With().Str(constants.PackageStrHelper, packageName).Logger()

	postDownstream := post.New(e)
	if err := postDownstream.Connect(); err != nil {
		logger.Err(err).Msg("failed to connect to post rpc server")
	}

	chatDownstream := chats.New(e)
	if err := chatDownstream.Connect(); err != nil {
		logger.Err(err).Msg("failed to connect to chat rpc server")
	}

	mediaDownstream := media.New(e)
	if err := mediaDownstream.Connect(); err != nil {
		logger.Err(err).Msg("failed to connect to media rpc server")
	}

	accountDownstream := account.New(e)
	if err := accountDownstream.Connect(); err != nil {
		logger.Err(err).Msg("failed to connect to account rpc server")
	}

	callsDownstream := calls.New(e)
	if err := callsDownstream.Connect(); err != nil {
		logger.Err(err).Msg("failed to connect to calls rpc server")
	}

	return &Downstream{
		logger: &logger,
		env:    e,
		Post:   postDownstream,
		Chat:   chatDownstream,
		Call:   callsDownstream,
		Media:  mediaDownstream,
	}
}
