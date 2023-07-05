package generator

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity"
	"github.com/kingmidas74/gonesis-engine/internal/domain/errors"
)

type EmptyGenerator struct{}

func (g EmptyGenerator) Generate(width, height int) (maze []contracts.Cell, err error) {
	if width <= 0 || height <= 0 {
		return make([]contracts.Cell, 0), errors.ErrMazeSizeIncorrect
	}

	maze = make([]contracts.Cell, width*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := entity.NewCell(x, y)
			c.SetWestWall(false)
			c.SetNorthWall(false)
			c.SetEastWall(false)
			c.SetSouthWall(false)
			maze[y*width+x] = c
		}
	}

	return maze, nil
}
