package wasm

import (
	"github.com/kingmidas74/gonesis-engine/internal/service/world"
)

type Params struct {
	WorldService world.Service
}

type Handler struct {
	worldService world.Service
}

func New(params Params) (*Handler, error) {
	return &Handler{
		worldService: params.WorldService,
	}, nil
}
