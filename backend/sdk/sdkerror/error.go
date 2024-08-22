package sdkerror

import "errors"

var (
	ErrNotEnoughArguments   = errors.New("not enough arguments")
	ErrUnableToConnectToDB  = errors.New("unable to connect to the DB")
	ErrDuplicateKey         = errors.New("duplicate key")
	ErrRecordCreation       = errors.New("record creation failed")
	ErrRecordNotFound       = errors.New("record not found")
	ErrFailedToDeleteRecord = errors.New("failed to delete record")
	ErrFailedToUpdateRecord = errors.New("failed to update record")

	ErrSomethingWentWrong = errors.New("something went wrong")
	//
)
