package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/wapp/backend/gateway/app"
	"github.com/dark-vinci/wapp/backend/gateway/env"
	"github.com/dark-vinci/wapp/backend/gateway/handlers/api"
	"github.com/dark-vinci/wapp/backend/gateway/handlers/websocket"
	"github.com/dark-vinci/wapp/backend/gateway/middleware"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

const packageName = "gateway.handlers"

type Handler struct {
	log        *zerolog.Logger
	env        *env.Environment
	app        *app.Operations
	middleware *middleware.Middleware
	engine     *gin.Engine
}

func New(e *env.Environment, log zerolog.Logger) *Handler {
	a := app.New(&log, e)

	r := gin.Default()
	mw := middleware.New(log, e, a)

	logger := log.With().Str(constants.PackageStrHelper, packageName).Logger()

	return &Handler{
		env:        e,
		log:        &logger,
		app:        &a,
		engine:     r,
		middleware: mw,
	}
}

func (h *Handler) Build(ctx context.Context) {
	gin.ForceConsoleColor()

	h.engine.Use(h.middleware.Cors())

	// build endpoints for REST API
	api.Build(h.engine.Group("/api"))

	// build endpoints for websocket
	websocket.New(ctx, *h.log, h.env, h.engine.Group("/socket"), h.app)
}

func (h *Handler) GetEngine() *gin.Engine {
	return h.engine
}
