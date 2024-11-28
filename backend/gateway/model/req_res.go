package model

import "github.com/google/uuid"

type (
	Error struct {
		PrivateMessage string `json:"privateMessage"`
		PublicMessage  string `json:"publicMessage"`
	}

	AppResponse[T any] struct {
		Data       *T
		Error      *Error
		RequestID  uuid.UUID
		StatusCode uint
	}
)

func NewSuccessResponse[T any](data *T, requestID uuid.UUID, statusCode int) *AppResponse[T] {
	return &AppResponse[T]{
		Data:       data,
		Error:      nil,
		RequestID:  requestID,
		StatusCode: uint(statusCode),
	}
}

func NewError[T any](privateMessage string, publicMessage string, requestID uuid.UUID, statusCode int) *AppResponse[T] {
	return &AppResponse[T]{
		Data:       nil,
		Error:      &Error{privateMessage, publicMessage},
		RequestID:  requestID,
		StatusCode: uint(statusCode),
	}
}
