package handlers

import (
	"fmt"
	socket "github.com/googollee/go-socket.io"
)

func (h *Handler) chat(s socket.Conn, msg string) string {
	s.SetContext(msg)

	return fmt.Sprintf("recv %s", msg)
}
