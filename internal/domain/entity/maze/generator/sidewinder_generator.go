package generator

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/errors"
	"github.com/kingmidas74/gonesis-engine/internal/util"
)

type SidewinderGenerator struct {
	gridGenerator GridGenerator
}

func (g SidewinderGenerator) Generate(width, height int) (maze []bool, err error) {
	if width <= 0 || height <= 0 {
		return nil, errors.ErrMazeSizeIncorrect
	}

	maze, err = g.gridGenerator.Generate(width, height)
	if err != nil {
		return nil, err
	}

	for y := 0; y < height; y = y + 2 {
		runset := make([]int, 0)
		for x := 0; x < width; x = x + 2 {

			if y == 0 {
				if x+1 < width {
					maze[y*width+x+1] = true
				}
				continue
			}

			runset = append(runset, x)

			direction := util.RandomIntBetween(0, 1)

			if direction == 1 && x+1 < width {
				maze[y*width+x+1] = true
				continue
			}

			randX := runset[util.RandomIntBetween(0, len(runset)-1)]
			maze[(y-1)*width+randX] = true
			runset = make([]int, 0)
		}
	}

	maze[0] = true

	return maze, nil
}
