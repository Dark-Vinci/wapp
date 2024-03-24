package model

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetStringPointer(v string) *string {
	return &v
}

type (
	ErrorData struct {
		ID             uuid.UUID `json:"id"`
		PublicMessage  string    `json:"publicMessage"`
		PrivateMessage string    `json:"privateMessage"`
		Handler        string    `json:"handler"`
	}

	GenericResponse[D any] struct {
		Data    *D         `json:"data"`
		Error   *ErrorData `json:"error"`
		Status  int        `json:"status"`
		Message *string    `json:"message"`
	}
)

func BuildResponse[D any](statusCode int, data *D, error *ErrorData, message *string) *GenericResponse[D] {
	return &GenericResponse[D]{
		Data:    data,
		Error:   error,
		Status:  statusCode,
		Message: message,
	}
}

func BuildErrorResponse[D any](c *gin.Context, statusCode int, error ErrorData) {
	c.JSON(statusCode, BuildResponse[D](
		statusCode,
		nil,
		&error,
		GetStringPointer("something went wrong"),
	))

	c.Abort()
}

func OkResponse[D any](c *gin.Context, statusCode int, data *D, message string) {
	c.JSON(statusCode, BuildResponse[D](
		statusCode,
		data,
		nil,
		GetStringPointer(message),
	))
	c.Abort()
}
