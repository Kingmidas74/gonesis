package terrain

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/contract"
	"github.com/kingmidas74/gonesis-engine/pkg/mod"
)

func (t *Terrain) Cells() []contract.Cell {
	return t.cells
}

func (t *Terrain) Cell(x, y int) contract.Cell {
	return t.cells[t.getCellIndex(x, y)]
}

func (t *Terrain) GetNeighbor(x, y, direction int) contract.Cell {
	coords := t.Topology.GetNeighbor(x, y, direction)
	return t.Cell(coords.X(), coords.Y())
}

func (t *Terrain) GetNeighbors(x, y int) []contract.Cell {
	coords := t.Topology.GetNeighbors(x, y)
	result := make([]contract.Cell, 0, len(coords))
	for _, c := range coords {
		result = append(result, t.Cell(c.X(), c.Y()))
	}
	return result
}

func (t *Terrain) CanMoveTo(from, to contract.Cell) bool {
	return t.Topology.CanMoveTo(from, to, t)
}

func (t *Terrain) Width() int {
	return t.width
}

func (t *Terrain) Height() int {
	return t.height
}

func (t *Terrain) EmptyCells() []contract.Cell {
	result := make([]contract.Cell, 0)
	for _, cell := range t.cells {
		if cell.IsEmpty() {
			result = append(result, cell)
		}
	}
	return result
}

func (t *Terrain) transformX(x int) int {
	return mod.ModLikePython(x, t.width)
}

func (t *Terrain) transformY(y int) int {
	return mod.ModLikePython(y, t.height)
}

func (t *Terrain) getCellIndex(x, y int) int {
	tx := t.transformX(x)
	ty := t.transformY(y)
	return ty*t.width + tx
}
