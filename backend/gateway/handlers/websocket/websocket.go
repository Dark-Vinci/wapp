package websocket

import (
	"context"
	"github.com/dark-vinci/wapp/backend/gateway/app"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/wapp/backend/gateway/env"
)

const packageName = "handler.websocket"

func New(ctx context.Context, log zerolog.Logger, e *env.Environment, r *gin.RouterGroup, app *app.Operations) {
	logger := log.With().Str("packageName", packageName).Logger()
	hub := NewHub(ctx, log, e)

	ws := WebSocket{
		upgrade: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		Hub:    hub,
		logger: logger,
	}

	ws.Build(r)
}

type WebSocket struct {
	upgrade websocket.Upgrader
	Hub     *Hub
	logger  zerolog.Logger
	//app     *app.Operations
}

func (ws *WebSocket) Build(endpoint *gin.RouterGroup) {
	// start the HUB
	ws.Hub.Start()

	endpoint.GET("/ws", func(c *gin.Context) {
		ws.serveWs(ws.Hub, c.Writer, c.Request)
	})
}

func (ws *WebSocket) serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	// AUTHENTICATE before doing this

	conn, err := ws.upgrade.Upgrade(w, r, nil)
	if err != nil {
		ws.logger.Err(err).Msgf("error upgrading ws: %v", err)
		return
	}

	client := NewClient(hub, conn, ws.logger)

	// write to user
	go client.WritePump()
	// read from user
	go client.ReadPump()
}
