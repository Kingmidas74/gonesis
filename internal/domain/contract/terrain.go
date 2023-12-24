package contract

type Terrain interface {
	Cell(x, y int) Cell
	Cells() []Cell
	EmptyCells() []Cell

	GetNeighbor(x, y, direction int) Cell
	GetNeighbors(x, y int) []Cell

	Width() int
	Height() int

	CanMoveTo(from, to Cell) bool
}
