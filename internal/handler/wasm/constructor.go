package wasm

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/service/world"
)

type Handler struct {
	worldService world.Service
}

func New() (*Handler, error) {
	return &Handler{
		worldService: world.New(configuration.NewConfiguration()),
	}, nil
}
