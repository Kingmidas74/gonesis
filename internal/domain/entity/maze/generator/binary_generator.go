package generator

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/errors"
	"github.com/kingmidas74/gonesis-engine/internal/util"
)

type BinaryGenerator struct {
	gridGenerator GridGenerator
}

func (g BinaryGenerator) Generate(width, height int) (maze []bool, err error) {
	if width <= 0 || height <= 0 {
		return make([]bool, 0), errors.ErrMazeSizeIncorrect
	}

	maze, err = g.gridGenerator.Generate(width, height)
	if err != nil {
		return nil, err
	}

	for y := 0; y < height; y = y + 2 {
		for x := 0; x < width; x = x + 2 {

			if y == 0 {
				if x+1 < width {
					maze[y*width+x+1] = true
				}
				continue
			}

			direction := util.RandomIntBetween(0, 1)
			if direction == 0 {
				maze[(y-1)*width+x] = true
				continue
			}

			if x+1 >= width {
				continue
			}

			maze[y*width+x+1] = true
		}
	}

	maze[0] = true

	return maze, nil
}
