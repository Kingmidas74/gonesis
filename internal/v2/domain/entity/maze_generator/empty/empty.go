package empty

import (
	"errors"

	"github.com/kingmidas74/gonesis-engine/internal/domain/entity"
)

type Cell interface {
	SetWestWall(bool)
	SetNorthWall(bool)
	SetEastWall(bool)
	SetSouthWall(bool)
}

type Generator struct{}

func New() *Generator {
	return &Generator{}
}

func (g Generator) Generate(width, height int) (maze []Cell, err error) {
	if width <= 0 || height <= 0 {
		return make([]Cell, 0), ErrMazeSizeIncorrect
	}

	maze = make([]Cell, width*height)

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

var ErrMazeSizeIncorrect = errors.New("can't create maze")
