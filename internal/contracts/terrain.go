package contracts

import "github.com/kingmidas74/gonesis-engine/internal/domain/enum"

type Terrain interface {
	GetNeighbor(x, y int, direction int) Cell
	Cell(x, y int) Cell
	Cells() []Cell
	Width() int
	Height() int
	SetCellType(x, y int, cell enum.CellType)
	EmptyCells() []Cell
}
