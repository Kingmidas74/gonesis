package moore

import (
	"github.com/kingmidas74/gonesis-engine/internal/util"
	"github.com/kingmidas74/gonesis-engine/internal/v2/domain/entity/coordinates"
)

type Coords interface {
	X() int
	SetX(x int)

	Y() int
	SetY(y int)
}

type Cell interface {
	X() int
	Y() int

	NorthWall() bool
	EastWall() bool
	SouthWall() bool
	WestWall() bool

	IsEmpty() bool
}

type Terrain interface {
	Width() int
	Height() int
	Cell(x, y int) Cell
}

type Topology struct {
}

func (t Topology) GetNeighbor(x, y int, direction int) Coords {
	multiples := t.getCoordsMultiples()
	d := Direction(util.ModLikePython(direction, len(t.getCoordsMultiples())))

	return coordinates.New(x+multiples[d][0], y+multiples[d][1])
}

func (t Topology) GetNeighbors(x, y int) []Coords {
	coordsMultiples := t.getCoordsMultiples()
	result := make([]Coords, len(coordsMultiples))

	for i, coords := range coordsMultiples {
		result[i] = coordinates.New(x+coords[0], y+coords[1])
	}

	return result
}

func (t Topology) CanMoveTo(currentCell, targetCell Cell, terra Terrain) bool {
	deltaX := targetCell.X() - currentCell.X()
	deltaY := targetCell.Y() - currentCell.Y()

	if deltaX > 1 {
		deltaX = deltaX - terra.Width()
	}
	if deltaX < -1 {
		deltaX = deltaX + terra.Width()
	}
	if deltaY > 1 {
		deltaY = deltaY - terra.Height()
	}
	if deltaY < -1 {
		deltaY = deltaY + terra.Height()
	}

	switch {
	case deltaX == 1 && deltaY == 0: // Moving East
		return !currentCell.EastWall()
	case deltaX == -1 && deltaY == 0: // Moving West
		return !currentCell.WestWall()
	case deltaY == 1 && deltaX == 0: // Moving South
		return !currentCell.SouthWall()
	case deltaY == -1 && deltaX == 0: // Moving North
		return !currentCell.NorthWall()
	case deltaX == 1 && deltaY == 1: // Moving Southeast
		return (terra.Cell(currentCell.X()+1, currentCell.Y()).IsEmpty() && !targetCell.NorthWall() && !currentCell.EastWall()) || (terra.Cell(currentCell.X(), currentCell.Y()+1).IsEmpty() && !targetCell.WestWall() && !currentCell.SouthWall())
	case deltaX == 1 && deltaY == -1: // Moving Northeast
		return (terra.Cell(currentCell.X()+1, currentCell.Y()).IsEmpty() && !targetCell.SouthWall() && !currentCell.EastWall()) || (terra.Cell(currentCell.X(), currentCell.Y()-1).IsEmpty() && !targetCell.WestWall() && !currentCell.NorthWall())
	case deltaX == -1 && deltaY == 1: // Moving Southwest
		return (terra.Cell(currentCell.X()-1, currentCell.Y()).IsEmpty() && !targetCell.SouthWall() && !currentCell.WestWall()) || (terra.Cell(currentCell.X(), currentCell.Y()+1).IsEmpty() && !targetCell.EastWall() && !currentCell.SouthWall())
	case deltaX == -1 && deltaY == -1: // Moving Northwest
		return (terra.Cell(currentCell.X()-1, currentCell.Y()).IsEmpty() && !targetCell.EastWall() && !currentCell.NorthWall()) || (terra.Cell(currentCell.X(), currentCell.Y()-1).IsEmpty() && !targetCell.SouthWall() && !currentCell.WestWall())
	default:
		return false // Target cell is the current cell
	}
}

func (t Topology) getCoordsMultiples() map[Direction][2]int {
	multiples := make(map[Direction][2]int)
	multiples[DirectionUp] = [2]int{0, -1}
	multiples[DirectionUpRight] = [2]int{1, -1}
	multiples[DirectionRight] = [2]int{1, 0}
	multiples[DirectionRightDown] = [2]int{1, 1}
	multiples[DirectionDown] = [2]int{0, 1}
	multiples[DirectionDownLeft] = [2]int{-1, 1}
	multiples[DirectionLeft] = [2]int{-1, 0}
	multiples[DirectionLeftUp] = [2]int{-1, -1}
	return multiples
}
