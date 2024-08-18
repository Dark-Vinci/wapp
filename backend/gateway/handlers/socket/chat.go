package socket

import (
	"fmt"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	socket "github.com/googollee/go-socket.io"
)

func (a *AppSocket) chat(s socket.Conn, msg string) string {
	s.SetContext(msg)

	log := a.log.With().
		Str(constants.PackageStrHelper, packageName).
		Str(constants.MethodStrHelper, "socket.chat").
		Logger()

	log.Info().Msg("Got a chat message")

	//validate
	// save to db

	//broadcast

	//a.server.

	//a.
	s.Emit("MESSAGE", msg)

	return fmt.Sprintf("recv %s", msg)
}
