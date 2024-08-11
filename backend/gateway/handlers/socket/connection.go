package socket

import (
	socket "github.com/googollee/go-socket.io"
)

func (a *AppSocket) connect(s socket.Conn) error {
	s.SetContext("")

	a.log.Debug().Msgf("connected: %s", s.ID())

	return nil
}

func (a *AppSocket) disconnect(_ socket.Conn, msg string) {
	a.log.Debug().Msgf("closed %s", msg)
}

func (a *AppSocket) error(_ socket.Conn, e error) {
	a.log.Debug().Msgf("meet error %s:", e)
}
