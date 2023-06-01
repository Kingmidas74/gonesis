package generator

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/errors"
)

type EmptyGenerator struct {
	gridGenerator GridGenerator
}

func (g EmptyGenerator) Generate(width, height int) (maze []bool, err error) {
	if width <= 0 || height <= 0 {
		return make([]bool, 0), errors.ErrMazeSizeIncorrect
	}

	maze = make([]bool, width*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			maze[y*width+x] = true
		}
	}

	return maze, nil
}
