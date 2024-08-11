package socket

import (
	"fmt"
	socket "github.com/googollee/go-socket.io"
)

func (a *AppSocket) chat(s socket.Conn, msg string) string {
	s.SetContext(msg)

	return fmt.Sprintf("recv %s", msg)
}
