package generator

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity"
	"github.com/kingmidas74/gonesis-engine/internal/domain/errors"
	"github.com/kingmidas74/gonesis-engine/internal/util"
)

type BinaryGenerator struct{}

func (g BinaryGenerator) Generate(width, height int) (maze []contracts.Cell, err error) {
	if width <= 0 || height <= 0 {
		return nil, errors.ErrMazeSizeIncorrect
	}

	// Initialize maze with all walls present
	maze = make([]contracts.Cell, width*height)
	for i := range maze {
		maze[i] = entity.NewCell(0, 0)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			maze[y*width+x].SetX(x)
			maze[y*width+x].SetY(y)

			if y == 0 {
				if x+1 < width {
					maze[y*width+x].SetEastWall(false)
					maze[y*width+x+1].SetWestWall(false)
					continue
				}
			}

			direction := util.RandomIntBetween(0, 1)
			if direction == 0 {
				if y > 0 {
					maze[y*width+x].SetNorthWall(false)
					maze[(y-1)*width+x].SetSouthWall(false)
				}
				continue
			}

			if x+1 >= width {
				continue
			}

			maze[y*width+x].SetEastWall(false)
			maze[y*width+x+1].SetWestWall(false)
		}
	}

	return maze, nil
}
