package socket

import (
	"github.com/rs/zerolog"

	"github.com/dark-vinci/wapp/backend/gateway/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	socket "github.com/googollee/go-socket.io"
)

const packageName = "gateway.handler.socket"

type AppSocket struct {
	env          *env.Environment
	log          *zerolog.Logger
	server       *socket.Server
	connectionID []string
}

const chatNamespace = "CHAT"
const defaultNamespace = ""
const groupNamespace = "GROUP"
const channelNamespace = "FANS"
const statusNamespace = "STATUS"

func Server(e *env.Environment, logger zerolog.Logger) *socket.Server {
	log := logger.With().
		Str(constants.PackageStrHelper, packageName).
		Logger()

	server := socket.NewServer(nil)

	if _, err := server.Adapter(&socket.RedisAdapterOptions{}); err != nil {
		log.Err(err).Msg("Failed to create redis adapter")
		panic(err)
	}

	a := AppSocket{
		env:          e,
		log:          &log,
		server:       server,
		connectionID: []string{},
	}

	server.OnConnect(defaultNamespace, a.connect)
	server.OnDisconnect(defaultNamespace, a.disconnect)
	server.OnError(defaultNamespace, a.error)

	server.OnEvent(chatNamespace, "SEND_MESSAGE", a.chat)
	server.OnEvent(chatNamespace, "TYPING", a.chat)

	server.OnEvent(groupNamespace, "msg", a.chat)
	server.OnEvent(groupNamespace, "TYPING", a.chat)

	server.OnEvent(channelNamespace, "msg", a.chat)

	server.OnEvent(statusNamespace, "msg", a.chat)
	server.OnEvent(defaultNamespace, "notice", a.chat)
	server.OnEvent(defaultNamespace, "bye", a.chat)

	return server
}
