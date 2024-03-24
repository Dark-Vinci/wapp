package error

import "errors"

var (
	ErrNotFound            = errors.New("requested resource was not found")
	ErrServiceNotAvailable = errors.New("the service in question is not online")
	ErrSomethingWentWrong  = errors.New("something went wrong try again later")
)
