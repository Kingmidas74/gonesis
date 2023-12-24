package maze_generator_collection

import (
	"errors"
	"github.com/kingmidas74/gonesis-engine/internal/domain/contract"
)

type MazeGeneratorCollection struct {
	mazeGenerators map[contract.MazeType]contract.MazeGenerator
}

func New() *MazeGeneratorCollection {
	return &MazeGeneratorCollection{
		mazeGenerators: map[contract.MazeType]contract.MazeGenerator{},
	}
}

var ErrMazeGenerationTypeIsNotSupported = errors.New("maze generation type is not supported")
