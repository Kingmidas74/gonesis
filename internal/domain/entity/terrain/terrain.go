package terrain

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"github.com/kingmidas74/gonesis-engine/internal/util"
)

type Terrain[T contracts.Topology] struct {
	cells  []contracts.Cell
	width  int
	height int
}

func NewTerrain[T contracts.Topology](maze contracts.Maze) contracts.Terrain {
	terra := &Terrain[T]{
		width:  maze.Width(),
		height: maze.Height(),
		cells:  maze.Content(),
	}

	for y := 0; y < maze.Height(); y++ {
		for x := 0; x < maze.Width(); x++ {
			cellIndex := terra.getCellIndex(x, y)
			terra.cells[cellIndex].IncreaseEnergy(maze.Height() - y)
		}
	}

	return terra
}

func (t *Terrain[T]) transformX(x int) int {
	return util.ModLikePython(x, t.width)
}

func (t *Terrain[T]) transformY(y int) int {
	return util.ModLikePython(y, t.height)
}

func (t *Terrain[T]) getCellIndex(x, y int) int {
	tx := t.transformX(x)
	ty := t.transformY(y)
	return ty*t.width + tx
}

func (t *Terrain[T]) Cell(x, y int) contracts.Cell {
	return t.cells[t.getCellIndex(x, y)]
}

func (t *Terrain[T]) Cells() []contracts.Cell {
	return t.cells
}

func (t *Terrain[T]) Width() int {
	return t.width
}

func (t *Terrain[T]) Height() int {
	return t.height
}

func (t *Terrain[T]) GetNeighbor(x, y int, direction int) contracts.Cell {
	coords := (*new(T)).GetNeighbor(x, y, direction)
	return t.Cell(coords.X(), coords.Y())
}

func (t *Terrain[T]) CanMoveTo(currentCell, targetCell contracts.Cell) bool {
	return targetCell.IsEmpty() && (*new(T)).CanMoveTo(currentCell, targetCell, t)
}

func (t *Terrain[T]) GetNeighbors(x, y int) []contracts.Cell {
	coords := (*new(T)).GetNeighbors(x, y)
	result := make([]contracts.Cell, 0, len(coords))
	for _, c := range coords {
		result = append(result, t.Cell(c.X(), c.Y()))
	}
	return result
}

func (t *Terrain[T]) SetCellType(x, y int, cellType enum.CellType) {
	t.cells[t.getCellIndex(x, y)].SetCellType(cellType)
}

func (t *Terrain[T]) EmptyCells() []contracts.Cell {
	result := make([]contracts.Cell, 0)
	for _, c := range t.Cells() {
		if c.IsEmpty() {
			result = append(result, c)
		}
	}
	return result
}
