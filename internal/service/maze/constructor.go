package maze

import "github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze"

type Service[K maze.Generator] struct{}

func NewMazeService[K maze.Generator]() *Service[K] {
	return &Service[K]{}
}
