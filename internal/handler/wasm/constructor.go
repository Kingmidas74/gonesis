package wasm

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/service/game"
)

type Handler struct {
	gameService *game.Service
}

func New() (*Handler, error) {
	return &Handler{
		gameService: game.New(configuration.NewConfiguration()),
	}, nil
}
