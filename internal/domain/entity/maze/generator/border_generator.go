package generator

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/errors"
)

type BorderGenerator struct {
	gridGenerator GridGenerator
}

func (g BorderGenerator) Generate(width, height int) (maze []bool, err error) {
	if width <= 0 || height <= 0 {
		return make([]bool, 0), errors.MAZE_SIZE_INCORRECT
	}

	maze = make([]bool, width*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			maze[y*width+x] = !((y == 0) || (x == 0) || (y == height-1) || (x == width-1))
		}
	}

	return maze, nil
}
