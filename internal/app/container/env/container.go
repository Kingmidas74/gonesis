package env

import (
	"go.uber.org/fx"

	"github.com/kingmidas74/gonesis-engine/internal/env"
)

func New() fx.Option {
	return fx.Provide(
		env.New,
	)
}
