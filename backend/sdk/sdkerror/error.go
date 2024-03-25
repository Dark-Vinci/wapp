package sdkerror

import "errors"

var (
	ErrNotEnoughArguments  = errors.New("not enough arguments")
	ErrUnableToConnectToDB = errors.New("unable to connect to the DB")
)
