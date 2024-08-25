package downstream

import (
	"github.com/rs/zerolog"

	"github.com/dark-vinci/wapp/backend/gateway/env"
)

const packageName = "gateway.downstream"

type Downstream struct{}

func New(z *zerolog.Logger, e *env.Environment) *Downstream {
	return &Downstream{}
}
