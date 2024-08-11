package handlers

import (
	//"github.com/dark-vinci/wapp/backend/gateway/handlers/socket"
	//"github.com/dark-vinci/wapp/backend/gateway/handlers/socket"
	"net/http"

	"github.com/gin-gonic/gin"
	socket "github.com/googollee/go-socket.io"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/wapp/backend/gateway/app"
	"github.com/dark-vinci/wapp/backend/gateway/env"
	appSocket "github.com/dark-vinci/wapp/backend/gateway/handlers/socket"
)

const packageName = "gateway.handlers"

type Handler struct {
	log          *zerolog.Logger
	env          *env.Environment
	app          *app.Operations
	socketServer *socket.Server
	middleware   *string
	engine       *gin.Engine
}

func New(e *env.Environment, logger zerolog.Logger) *Handler {
	a := app.New()
	s := appSocket.Server(e, logger)

	r := gin.Default()

	return &Handler{
		env:          e,
		log:          &logger,
		app:          &a,
		socketServer: s,
		engine:       r,
	}
}

func (h *Handler) GetEngine() *gin.Engine {
	return h.engine
}

func (h *Handler) Build() {
	gin.ForceConsoleColor()

	h.engine.Use(ginMiddleware(h.env.FrontEndURL))

	apiGroup := h.engine.Group("/api")

	apiGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name":     "tomato",
			"response": "200",
		})
	})

	//socket IO
	h.engine.GET("/socket.io/*any", abc, gin.WrapH(h.socketServer))
	h.engine.POST("/socket.io/*any", abc, gin.WrapH(h.socketServer))
	h.engine.StaticFS("/public", http.Dir("../asset"))
}

func abc(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name":     "tomato",
		"response": "200",
	})
}

func ginMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}
