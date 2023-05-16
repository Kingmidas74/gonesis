package terrain

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity"
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
		cells:  make([]contracts.Cell, maze.Width()*maze.Height()),
	}

	for y := 0; y < maze.Height(); y++ {
		for x := 0; x < maze.Width(); x++ {
			cellIndex := terra.getCellIndex(x, y)
			if maze.Content()[cellIndex] {
				terra.cells[cellIndex] = entity.NewCell(x, y, enum.CellTypeEmpty)
			} else {
				terra.cells[cellIndex] = entity.NewCell(x, y, enum.CellTypeObstacle)
			}
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
	x = t.transformX(x)
	y = t.transformY(y)
	return y*t.width + x
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

func (t *Terrain[T]) SetCellType(x, y int, cellType enum.CellType) {
	t.cells[t.getCellIndex(x, y)].SetCellType(cellType)
}

func (t *Terrain[T]) EmptyCells() []contracts.Cell {
	result := make([]contracts.Cell, 0)
	for _, c := range t.Cells() {
		if c.CellType() == enum.CellTypeEmpty {
			result = append(result, c)
		}
	}
	return result
}
