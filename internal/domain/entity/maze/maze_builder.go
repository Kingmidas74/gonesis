package maze

import "github.com/kingmidas74/gonesis-engine/internal/contracts"

type Builder[G contracts.MazeGenerator] struct {
	width       int
	height      int
	firstFilled bool

	maze Maze
}

func NewMazeBuilder[G contracts.MazeGenerator]() contracts.MazeBuilder[G] {
	return &Builder[G]{}
}

func (b *Builder[G]) SetWidth(width int) contracts.MazeBuilder[G] {
	b.width = width

	return b
}

func (b *Builder[G]) SetHeight(height int) contracts.MazeBuilder[G] {
	b.height = height

	return b
}

func (b *Builder[G]) FirstFilled(flag bool) contracts.MazeBuilder[G] {
	b.firstFilled = flag

	return b
}

func (b *Builder[G]) Build() (contracts.Maze, error) {
	g := *new(G)

	maze, err := g.Generate(b.width, b.height)
	if err != nil {
		return nil, err
	}

	return newMaze(b.width, b.height, maze), nil
}
