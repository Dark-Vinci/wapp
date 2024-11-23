package middleware

import (
	"github.com/rs/zerolog"

	"github.com/dark-vinci/wapp/backend/gateway/app"
	"github.com/dark-vinci/wapp/backend/gateway/env"
)

const packageName = "gateway.middleware"

type Middleware struct {
	logger zerolog.Logger
	env    *env.Environment
	app    app.Operations
}

func New(l zerolog.Logger, e *env.Environment, a app.Operations) *Middleware {
	logger := l.With().Str("MODULE", packageName).Logger()

	return &Middleware{
		logger: logger,
		env:    e,
		app:    a,
	}
}
