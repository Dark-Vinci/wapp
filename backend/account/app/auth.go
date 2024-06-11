package app

import (
	"context"
	"errors"
)

func (a *App) Signup(ctx context.Context, aB int) (*int, error) {
	a.logger.Debug().Msg("MORE LOGGER")

	return nil, errors.New("NEW ERROR")
}
