package websocket

import (
	"github.com/dark-vinci/wapp/backend/gateway/env"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
)

const packageName = "handler.websocket"

type WebSocket struct {
	upgrade websocket.Upgrader
	Hub     *Hub
	logger  zerolog.Logger
}

func New(log zerolog.Logger, e *env.Environment, r *gin.RouterGroup) {
	logger := log.With().Str("packageName", packageName).Logger()
	hub := NewHub(log, e)

	ww := WebSocket{
		upgrade: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		Hub:    hub,
		logger: logger,
	}

	ww.Build(r)
}

func (ws *WebSocket) Build(endpoint *gin.RouterGroup) {
	// start the HUB
	ws.Hub.Start()

	endpoint.GET("/ws", func(c *gin.Context) {
		ws.serveWs(ws.Hub, c.Writer, c.Request)
	})
}

func (ws *WebSocket) serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrade.Upgrade(w, r, nil)
	if err != nil {
		// todo: return an error
		log.Println(err)
	}

	client := NewClient(hub, conn, ws.logger)

	// write to user
	go client.WritePump()
	// read from user
	go client.ReadPump()
}
