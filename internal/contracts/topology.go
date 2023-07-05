package contracts

type Topology interface {
	GetNeighbor(x, y int, direction int) Coords
	GetNeighbors(x, y int) []Coords
	CanMoveTo(currentCell, targetCell Cell, terra Terrain) bool
}
