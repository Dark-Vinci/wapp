package posts

import (
	"github.com/dark-vinci/wapp/backend/gateway/middleware"
	"github.com/gin-gonic/gin"
)

type postApi struct {
	m *middleware.Middleware
}

func New(eng *gin.RouterGroup) {
	p := postApi{}

	postGroup := eng.Group("/post", p.m.Authenticate())

	postGroup.GET("/", func(context *gin.Context) {
	})
}
