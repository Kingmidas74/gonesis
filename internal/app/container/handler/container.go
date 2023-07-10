package handler

import (
	"go.uber.org/fx"

	"github.com/kingmidas74/gonesis-engine/internal/handler/http"
	"github.com/kingmidas74/gonesis-engine/internal/handler/http/middleware/no_cache"
)

func New() fx.Option {
	return fx.Provide(
		http.NewServer,

		no_cache.New,
	)
}
