package model

const packageName = "gateway.model"

type LoginRequest struct {
	PhoneNumber string `json:"username"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
