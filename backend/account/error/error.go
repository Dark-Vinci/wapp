package error

import "errors"

var (
	ErrAccountFailedToStart = errors.New("account service failed to start")
)
