package maze

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/service/maze"
)

type Handler[G contracts.MazeGenerator] struct {
	mazeService *maze.Service[G]
}

func New[G contracts.MazeGenerator]() *Handler[G] {
	return &Handler[G]{
		mazeService: maze.NewMazeService[G](),
	}
}
