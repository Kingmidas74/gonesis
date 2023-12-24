package contract

import "github.com/kingmidas74/gonesis-engine/pkg/maze/cell"

type MazeGenerator interface {
	Generate(width, height int) ([]cell.Cell, error)
}
