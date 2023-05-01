package maze

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/errors"
)

type Generator interface {
	Generate(width, height int, matrix []bool) []bool
}

type Builder[G Generator] struct {
	width       int
	height      int
	firstFilled bool

	maze Maze
}

func NewMazeBuilder[G Generator]() *Builder[G] {
	return &Builder[G]{}
}

func (b *Builder[G]) SetWidth(width int) *Builder[G] {
	b.width = width

	return b
}

func (b *Builder[G]) SetHeight(height int) *Builder[G] {
	b.height = height

	return b
}

func (b *Builder[G]) FirstFilled(flag bool) *Builder[G] {
	b.firstFilled = flag

	return b
}

func (b *Builder[G]) Build() (*Maze, error) {
	grid, err := b.generateGrid(b.width, b.height, b.firstFilled)
	if err != nil {
		return nil, err
	}

	g := *new(G)
	matrix := g.Generate(b.width, b.height, grid)

	return newMaze(b.width, b.height, matrix), nil
}

func (b *Builder[G]) generateGrid(width, height int, firstFilled bool) ([]bool, error) {
	if width <= 0 || height <= 0 {
		return make([]bool, 0), errors.MAZE_SIZE_INCORRECT
	}

	result := make([]bool, width*height)

	for y := 0; y < height; y = y + 2 {
		for x := 0; x < width; x = x + 2 {
			result[y*width+x] = true
		}
	}

	result[0] = firstFilled

	return result, nil
}
