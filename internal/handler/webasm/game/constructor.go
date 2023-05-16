package game

import (
	"github.com/kingmidas74/gonesis-engine/internal/service/game"
)

type Handler struct {
	gameService *game.Service
}

func New() *Handler {
	return &Handler{
		gameService: game.New(),
	}
}
