package account

import (
	"github.com/dark-vinci/wapp/backend/gateway/env"
	"github.com/dark-vinci/wapp/backend/gateway/middleware"
	"github.com/gin-gonic/gin"
)

type accountApi struct {
	m middleware.Middleware
	e *env.Environment
}

func New(eng *gin.RouterGroup) {
	a := accountApi{}

	accountGroup := eng.Group("/account", a.m.Authenticate())

	accountGroup.GET("/:id", func(ctx *gin.Context) {
	})
}
