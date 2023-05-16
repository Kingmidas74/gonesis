package generator

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/errors"
)

type GridGenerator struct {
}

func (g GridGenerator) Generate(width, height int) (maze []bool, err error) {
	if width <= 0 || height <= 0 {
		return make([]bool, 0), errors.ErrMazeSizeIncorrect
	}

	maze = make([]bool, width*height)
	for i := range maze {
		maze[i] = false
	}

	for y := 0; y < height; y = y + 2 {
		for x := 0; x < width; x = x + 2 {
			maze[y*width+x] = true
		}
	}

	maze[0] = true

	return maze, nil
}
