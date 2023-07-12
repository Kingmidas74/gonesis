package maze

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Service interface {
	Generate(mazeType enum.MazeType, width, height, requiredEmptyCells int) (contracts.Maze, error)
}
