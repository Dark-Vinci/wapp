package auth

import (
	"github.com/dark-vinci/wapp/backend/gateway/env"
	"github.com/dark-vinci/wapp/backend/gateway/middleware"
	"github.com/gin-gonic/gin"
)

type authApi struct {
	m *middleware.Middleware
	e *env.Environment
}

func New(eng *gin.RouterGroup) {
	auth := authApi{}

	auth.m.Cors()

	authGroup := eng.Group("/auth")

	authGroup.POST("/login", func(context *gin.Context) {
	})

	authGroup.POST("/sign-up-with-email", func(context *gin.Context) {
	})

	authGroup.POST("/sign-up-with-google", func(context *gin.Context) {
	})

	authGroup.POST("/verify-token", func(context *gin.Context) {
	})

	authGroup.POST("/get-token", func(context *gin.Context) {
	})
}
