package generator

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity"
	"github.com/kingmidas74/gonesis-engine/internal/domain/errors"
)

type GridGenerator struct {
}

func (g GridGenerator) Generate(width, height int) (maze []contracts.Cell, err error) {
	if width <= 0 || height <= 0 {
		return make([]contracts.Cell, 0), errors.ErrMazeSizeIncorrect
	}

	maze = make([]contracts.Cell, width*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := entity.NewCell(x, y)
			maze[y*width+x] = c
		}
	}

	return maze, nil
}
