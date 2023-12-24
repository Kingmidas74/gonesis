package terrain

import (
	contract "github.com/kingmidas74/gonesis-engine/internal/domain/contract"
)

type MazeInfo struct {
	Width  int
	Height int
	Cells  []contract.Cell
}

type Terrain struct {
	contract.Topology

	cells  []contract.Cell
	width  int
	height int
}

func New(maze MazeInfo, options ...func(terrain *Terrain)) *Terrain {
	t := &Terrain{
		cells:  maze.Cells,
		width:  maze.Width,
		height: maze.Height,
	}
	for _, o := range options {
		o(t)
	}
	return t
}

func WithTopology(topology contract.Topology) func(*Terrain) {
	return func(t *Terrain) {
		t.Topology = topology
	}
}
