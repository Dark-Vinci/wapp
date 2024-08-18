package socket

import (
	socket "github.com/googollee/go-socket.io"
)

func (a *AppSocket) connect(s socket.Conn) error {
	s.SetContext("")

	s.Emit("connect", s)

	//a.server.

	socketID := s.ID()

	a.connectionID = append(a.connectionID, socketID)

	//a.server.
	a.log.Debug().Msgf("connected: %s", s.ID())

	return nil
}

func (a *AppSocket) disconnect(_ socket.Conn, msg string) {
	a.log.Debug().Msgf("closed %s", msg)
}

func (a *AppSocket) error(_ socket.Conn, e error) {
	a.log.Debug().Msgf("meet error %s:", e)
}
