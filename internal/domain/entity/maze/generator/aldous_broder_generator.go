package generator

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity"
	"github.com/kingmidas74/gonesis-engine/internal/domain/errors"
	"github.com/kingmidas74/gonesis-engine/internal/util"
)

type AldousBroderGenerator struct{}

func (g AldousBroderGenerator) Generate(width, height int) (maze []contracts.Cell, err error) {
	if width <= 0 || height <= 0 {
		return nil, errors.ErrMazeSizeIncorrect
	}

	// Initialize maze with all walls present
	maze = make([]contracts.Cell, width*height)
	for i := range maze {
		maze[i] = entity.NewCell(0, 0)
	}

	directions := []struct {
		dx, dy int
	}{
		{0, -1}, // North
		{1, 0},  // East
		{0, 1},  // South
		{-1, 0}, // West
	}

	// Choose a random cell
	x, y := util.RandomIntBetween(0, width-1), util.RandomIntBetween(0, height-1)
	unvisited := width*height - 1

	for unvisited > 0 {
		direction := directions[util.RandomIntBetween(0, len(directions)-1)]
		newX, newY := x+direction.dx, y+direction.dy

		if newX >= 0 && newX < width && newY >= 0 && newY < height {
			index := newY*width + newX
			maze[index].SetX(newX)
			maze[index].SetY(newY)
			oldIndex := y*width + x

			hasAllWalls := maze[index].NorthWall() && maze[index].EastWall() && maze[index].SouthWall() && maze[index].WestWall()
			if hasAllWalls {
				switch {
				case direction.dx > 0:
					maze[oldIndex].SetEastWall(false)
					maze[index].SetWestWall(false)
				case direction.dx < 0:
					maze[oldIndex].SetWestWall(false)
					maze[index].SetEastWall(false)
				case direction.dy > 0:
					maze[oldIndex].SetSouthWall(false)
					maze[index].SetNorthWall(false)
				case direction.dy < 0:
					maze[oldIndex].SetNorthWall(false)
					maze[index].SetSouthWall(false)
				}
				unvisited--
			}
			x, y = newX, newY
		}
	}

	return maze, nil
}
