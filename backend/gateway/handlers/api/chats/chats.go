package chats

import (
	"github.com/dark-vinci/wapp/backend/gateway/env"
	"github.com/dark-vinci/wapp/backend/gateway/middleware"
	"github.com/gin-gonic/gin"
)

type chatApi struct {
	m *middleware.Middleware
	e *env.Environment
}

func New(eng *gin.RouterGroup) {
	c := chatApi{}

	chat := eng.Group("/chat", c.m.Authenticate())

	personalChat := chat.Group("/personal")
	personalChat.GET("/", func(context *gin.Context) {

	})

	groupChat := chat.Group("/group")
	groupChat.GET("/", func(context *gin.Context) {

	})

	channelChat := chat.Group("/channel")
	channelChat.GET("/", func(context *gin.Context) {

	})
}
