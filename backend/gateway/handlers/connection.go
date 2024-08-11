package handlers

import (
	socket "github.com/googollee/go-socket.io"
)

func (h *Handler) connect(s socket.Conn) error {
	s.SetContext("")

	h.log.Println("connected:", s.ID())

	return nil
}

func (h *Handler) disconnect(s socket.Conn, msg string) {
	h.log.Println("closed", msg)
}
