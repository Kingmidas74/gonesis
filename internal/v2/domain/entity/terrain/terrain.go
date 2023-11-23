package terrain

type Cell interface {
}

type Maze interface {
	Width() int
	Height() int
	Content() []Cell
}

type Topology interface {
}

type Terrain struct {
	Topology

	cells  []Cell
	width  int
	height int
}

func New(maze Maze, options ...func(terrain *Terrain)) *Terrain {
	t := &Terrain{
		cells:  maze.Content(),
		width:  maze.Width(),
		height: maze.Height(),
	}
	for _, o := range options {
		o(t)
	}
	return t
}

func WithTopology(topology Topology) func(*Terrain) {
	return func(t *Terrain) {
		t.Topology = topology
	}
}

func (t *Terrain) Cells() []Cell {
	return t.cells
}
