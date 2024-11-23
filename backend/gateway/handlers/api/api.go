package api

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/dark-vinci/wapp/backend/gateway/handlers/api/account"
	"github.com/dark-vinci/wapp/backend/gateway/handlers/api/auth"
	"github.com/dark-vinci/wapp/backend/gateway/handlers/api/chats"
	"github.com/dark-vinci/wapp/backend/gateway/handlers/api/posts"
)

func Build(r *gin.RouterGroup) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name":     "tomato",
			"response": "200",
		})
	})

	account.New(r)
	auth.New(r)
	chats.New(r)
	posts.New(r)
}
