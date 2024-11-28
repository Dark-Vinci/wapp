package auth

import (
	"github.com/dark-vinci/wapp/backend/gateway/app"
	"github.com/dark-vinci/wapp/backend/gateway/env"
	"github.com/dark-vinci/wapp/backend/gateway/middleware"
	"github.com/dark-vinci/wapp/backend/gateway/model"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
)

type authApi struct {
	m      *middleware.Middleware
	e      *env.Environment
	logger *zerolog.Logger
	app    app.Operations
}

func (a *authApi) login(ctx *gin.Context) {
	log := a.logger.With().Str("endpoint", ctx.FullPath()).Logger()

	var credential model.LoginRequest

	if err := ctx.ShouldBindJSON(&credential); err != nil {
		log.Err(err).Msg("fail to parse login request credentials")

		ctx.JSON(http.StatusOK, gin.H{"message": "User validation successful."})
		return
	}

	if err := model.Validate(credential); err != nil {

	}

	response, err := a.app.LoginToAccount(ctx, credential)
	if err != nil {
		log.Err(err).Msg("fail to login")
		return //todo' update
	}

	ctx.JSON(http.StatusOK, response) //todo;
}

func New(eng *gin.RouterGroup) {
	auth := authApi{}

	auth.m.Cors()

	authGroup := eng.Group("/auth")

	authGroup.POST("/login", func(ctx *gin.Context) {
		//type LL struct {
		//	PhoneNumber string `json:"username"`
		//	Password    string `json:"password"`
		//}
		//
		//var credential LL
		//var endpoint = ctx.FullPath()
		//
		//if err := ctx.ShouldBindJSON(&credential); err != nil {
		//	ctx.JSON(http.StatusOK, gin.H{"message": "User validation successful."})
		//}

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
