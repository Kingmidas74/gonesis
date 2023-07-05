package topology

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity"
	"github.com/kingmidas74/gonesis-engine/internal/util"
)

type MooreDirection uint8

const (
	MooreDirectionUp MooreDirection = iota
	MooreDirectionUpRight
	MooreDirectionRight
	MooreDirectionRightDown
	MooreDirectionDown
	MooreDirectionDownLeft
	MooreDirectionLeft
	MooreDirectionLeftUp
)

func (d MooreDirection) Value() uint8 {
	return uint8(d)
}

type MooreTopology struct {
}

func (t MooreTopology) GetNeighbor(x, y int, direction int) contracts.Coords {
	multiples := t.getCoordsMultiples()
	mooreDirection := MooreDirection(util.ModLikePython(direction, len(t.getCoordsMultiples())))

	return entity.NewCoords(x+multiples[mooreDirection][0], y+multiples[mooreDirection][1])
}

func (t MooreTopology) GetNeighbors(x, y int) []contracts.Coords {
	coordsMultiples := t.getCoordsMultiples()
	result := make([]contracts.Coords, len(coordsMultiples))

	for i, coords := range coordsMultiples {
		result[i] = entity.NewCoords(x+coords[0], y+coords[1])
	}

	return result
}

func (t MooreTopology) CanMoveTo(currentCell, targetCell contracts.Cell, terra contracts.Terrain) bool {
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

func (t MooreTopology) getCoordsMultiples() map[MooreDirection][2]int {
	multiples := make(map[MooreDirection][2]int)
	multiples[MooreDirectionUp] = [2]int{0, -1}
	multiples[MooreDirectionUpRight] = [2]int{1, -1}
	multiples[MooreDirectionRight] = [2]int{1, 0}
	multiples[MooreDirectionRightDown] = [2]int{1, 1}
	multiples[MooreDirectionDown] = [2]int{0, 1}
	multiples[MooreDirectionDownLeft] = [2]int{-1, 1}
	multiples[MooreDirectionLeft] = [2]int{-1, 0}
	multiples[MooreDirectionLeftUp] = [2]int{-1, -1}
	return multiples
}
