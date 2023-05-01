package maze

import (
	domain_maze "github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze"
	"github.com/kingmidas74/gonesis-engine/internal/service/maze"
)

type Handler[G domain_maze.Generator] struct {
	mazeService *maze.Service[G]
}

func New[G domain_maze.Generator]() *Handler[G] {
	return &Handler[G]{
		mazeService: maze.NewMazeService[G](),
	}
}
