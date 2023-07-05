package generator

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity"
	"github.com/kingmidas74/gonesis-engine/internal/domain/errors"
	"github.com/kingmidas74/gonesis-engine/internal/util"
)

type SidewinderGenerator struct{}

func (g SidewinderGenerator) Generate(width, height int) (maze []contracts.Cell, err error) {
	if width <= 0 || height <= 0 {
		return nil, errors.ErrMazeSizeIncorrect
	}

	// Initialize maze with all walls present
	maze = make([]contracts.Cell, width*height)
	for i := range maze {
		maze[i] = entity.NewCell(0, 0)
	}

	for y := 0; y < height; y++ {
		runset := make([]int, 0)
		for x := 0; x < width; x++ {
			index := y*width + x
			maze[index].SetX(x)
			maze[index].SetY(y)

			if y == 0 {
				if x+1 < width {
					maze[index].SetEastWall(false)
					maze[index+1].SetWestWall(false)
				}
				continue
			}

			runset = append(runset, x)

			direction := util.RandomIntBetween(0, 1)

			if direction == 1 && x+1 < width {
				maze[index].SetEastWall(false)
				maze[index+1].SetWestWall(false)
				continue
			}

			randX := runset[util.RandomIntBetween(0, len(runset)-1)]
			maze[y*width+randX].SetNorthWall(false)
			maze[(y-1)*width+randX].SetSouthWall(false)
			runset = make([]int, 0)
		}
	}

	return maze, nil
}
