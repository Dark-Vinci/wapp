package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//socket "github.com/googollee/go-socket.io"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/wapp/backend/gateway/app"
	"github.com/dark-vinci/wapp/backend/gateway/env"
	"github.com/dark-vinci/wapp/backend/gateway/handlers/api"
	"github.com/dark-vinci/wapp/backend/gateway/handlers/websocket"
	"github.com/dark-vinci/wapp/backend/gateway/middleware"
)

const packageName = "gateway.handlers"

type Handler struct {
	log *zerolog.Logger
	env *env.Environment
	app *app.Operations
	//socketServer *socket.Server
	middleware *middleware.Middleware
	engine     *gin.Engine
}

func New(e *env.Environment, log zerolog.Logger) *Handler {
	a := app.New()
	//s := appSocket.Server(e, log)

	r := gin.Default()
	mw := middleware.New(log, e, a)

	logger := log.With().Str("PACKAGE", packageName).Logger()

	return &Handler{
		env: e,
		log: &logger,
		app: &a,
		//socketServer: s,
		engine:     r,
		middleware: mw,
	}
}

func (h *Handler) GetEngine() *gin.Engine {
	return h.engine
}

func (h *Handler) Build() {
	gin.ForceConsoleColor()

	// cors middleware
	h.engine.Use(h.middleware.Cors())

	apiGroup := h.engine.Group("/api")

	// add no middleware
	apiGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name":     "tomato",
			"response": "200",
		})
	})

	// logged-in and non
	api.New(apiGroup)

	// user must be logged in
	ws := websocket.New(*h.log, h.env)
	ws.Build(h.engine.Group("/socket"))
}
