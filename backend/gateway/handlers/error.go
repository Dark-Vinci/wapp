package handlers

import socket "github.com/googollee/go-socket.io"

func (h *Handler) error(s socket.Conn, e error) {
	h.log.Println("meet error:", e)
}
