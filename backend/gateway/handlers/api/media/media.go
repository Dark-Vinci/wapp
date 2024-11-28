package media

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/dark-vinci/wapp/backend/gateway/app"
	"github.com/dark-vinci/wapp/backend/gateway/env"
	"github.com/dark-vinci/wapp/backend/gateway/middleware"
)

type mediaApi struct {
	m   *middleware.Middleware
	e   *env.Environment
	app app.Operations
}

const single = "single"

func New(eng *gin.RouterGroup) {
	c := mediaApi{}

	media := eng.Group("/media", c.m.Authenticate())

	media.POST("/one", c.UploadSingle)
	media.POST("/multiple", c.UploadMultiple)
}

func (media *mediaApi) UploadSingle(ctx *gin.Context) {
	file, err := ctx.FormFile(single)
	if err != nil {
		return
	}

	userUUID, _ := uuid.Parse("USER_ID")

	result, err := media.app.UploadSingleMedia(context.Background(), userUUID, file)
	if err != nil {

	}

	ctx.JSON(http.StatusOK, result)
}

func (media *mediaApi) UploadMultiple(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		fmt.Println(err, "unable to get file")
		return
	}

	files := form.File["upload[]"]

	result, err := media.app.UploadMultipleMedia(context.Background(), files)
	if err != nil {
		fmt.Println(err, "unable to get file")
	}

	ctx.JSON(200, result)
}
