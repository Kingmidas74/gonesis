package maze_generator_collection

import "github.com/kingmidas74/gonesis-engine/internal/domain/contract"

func (c *MazeGeneratorCollection) Register(mazeType contract.MazeType, mazeGenerator contract.MazeGenerator) {
	c.mazeGenerators[mazeType] = mazeGenerator
}

func (c *MazeGeneratorCollection) Get(mazeType contract.MazeType) (contract.MazeGenerator, error) {
	mazeGenerator, ok := c.mazeGenerators[mazeType]
	if !ok {
		return nil, ErrMazeGenerationTypeIsNotSupported
	}
	return mazeGenerator, nil
}
