package maze

type Generator interface {
	Generate(width, height int) (maze []bool, err error)
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
	g := *new(G)

	maze, err := g.Generate(b.width, b.height)
	if err != nil {
		return nil, err
	}

	return newMaze(b.width, b.height, maze), nil
}
