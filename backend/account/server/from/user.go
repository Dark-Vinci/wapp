package from

import (
	"github.com/dark-vinci/wapp/backend/sdk/grpc/account"
	"github.com/dark-vinci/wapp/backend/sdk/models"
)

func CreateUserResponse(reqID string, user models.User) *account.CreateUserResponse {
	response := new(account.CreateUserResponse)

	response.RequestId = reqID
	response.Email = user.Email
	response.FirstName = user.FirstName
	response.LastName = user.LastName

	return response
}

func CreateUserRequest(in *account.CreateUserRequest) models.User {
	response := models.User{
		Email:     in.Email,
		FirstName: in.FirstName,
		LastName:  in.LastName,
	}

	return response
}
