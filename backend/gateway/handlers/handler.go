package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	socket "github.com/googollee/go-socket.io"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/wapp/backend/gateway/app"
	"github.com/dark-vinci/wapp/backend/gateway/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

const packageName = "gateway.handlers"

const chatNamespace = "CHAT"
const defaultNamespace = ""
const groupNamespace = "GROUP"
const fansNamespace = "FANS"

type Handler struct {
	log *zerolog.Logger
	env *env.Environment
	app app.Operations
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

func Router(e *env.Environment, logger zerolog.Logger) *gin.Engine {
	log := logger.With().
		Str(constants.PackageStrHelper, packageName).
		Str(constants.FunctionNameHelper, "Router").
		Logger()

	r := gin.Default()

	a := app.New()

	h := Handler{
		log: &log,
		env: e,
		app: a,
	}

	server := socket.NewServer(nil)

	if _, err := server.Adapter(&socket.RedisAdapterOptions{}); err != nil {
		panic(err)
	}

	server.OnConnect(defaultNamespace, h.connect)
	server.OnDisconnect(defaultNamespace, h.disconnect)
	server.OnError(defaultNamespace, h.error)

	server.OnEvent(chatNamespace, "msg", h.chat)
	server.OnEvent(groupNamespace, "msg", h.chat)
	server.OnEvent(fansNamespace, "msg", h.chat)

	server.OnEvent(defaultNamespace, "notice", h.chat)
	server.OnEvent(defaultNamespace, "bye", h.chat)

	gin.ForceConsoleColor()

	r.Use(ginMiddleware(e.FrontEndURL))

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name":     "tomato",
			"response": "200",
		})
	})

	// auth, handler
	// secretes
	// images
	//

	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))
	r.StaticFS("/public", http.Dir("../asset"))

	return r
}
