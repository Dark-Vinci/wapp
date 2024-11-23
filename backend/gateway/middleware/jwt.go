package middleware

import (
	"time"
)

type Token struct {
	Value            string    `json:"value"`
	Refresh          string    `json:"refresh"`
	ValueExpiresAt   time.Time `json:"value_expires_at"`
	RefreshExpiresAt time.Time `json:"refresh_expires_at"`
}

func (m *Middleware) CreateToken() (*Token, error) {
	return &Token{}, nil
}

func (m *Middleware) ParseToken(token string) (*Token, error) {
	return &Token{}, nil
}

func (m *Middleware) ValidateRefreshToken(token string) (*Token, error) {
	return &Token{}, nil
}
