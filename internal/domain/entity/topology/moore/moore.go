package moore

import (
	contract "github.com/kingmidas74/gonesis-engine/internal/domain/contract"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/coordinate"
)

func (t Topology) GetNeighbor(x, y int, direction int) contract.Coordinate {
	neighbor := t.topology.GetNeighbor(x, y, direction)
	return coordinate.New(neighbor.X(), neighbor.Y())
}

func (t Topology) GetNeighbors(x, y int) []contract.Coordinate {
	neighbors := t.topology.GetNeighbors(x, y)

	result := make([]contract.Coordinate, len(neighbors))
	for i, neighbor := range neighbors {
		result[i] = coordinate.New(neighbor.X(), neighbor.Y())
	}

	return result
}

func (t Topology) CanMoveTo(currentCell, targetCell contract.Cell, terra contract.Terrain) bool {
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
