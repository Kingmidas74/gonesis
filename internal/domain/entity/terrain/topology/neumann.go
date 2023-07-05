package topology

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity"
	"github.com/kingmidas74/gonesis-engine/internal/util"
)

type NeumannDirection uint8

const (
	NeumannDirectionUp NeumannDirection = iota
	NeumannDirectionRight
	NeumannDirectionDown
	NeumannDirectionLeft
)

func (d NeumannDirection) Value() uint8 {
	return uint8(d)
}

type NeumannTopology struct {
}

func (t NeumannTopology) GetNeighbor(x, y int, direction int) contracts.Coords {
	multiples := t.getCoordsMultiples()
	neumannDirection := NeumannDirection(util.ModLikePython(direction, len(multiples)))

	return entity.NewCoords(x+multiples[neumannDirection][0], y+multiples[neumannDirection][1])
}

func (t NeumannTopology) GetNeighbors(x, y int) []contracts.Coords {
	coordsMultiples := t.getCoordsMultiples()
	result := make([]contracts.Coords, len(coordsMultiples))

	for i, coords := range coordsMultiples {
		result[i] = entity.NewCoords(x+coords[0], y+coords[1])
	}

	return result
}

func (t NeumannTopology) CanMoveTo(currentCell, targetCell contracts.Cell, terra contracts.Terrain) bool {
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
	default:
		return false // Target cell is the current cell
	}
}

func (t NeumannTopology) getCoordsMultiples() map[NeumannDirection][2]int {
	multiples := make(map[NeumannDirection][2]int)
	multiples[NeumannDirectionUp] = [2]int{0, -1}
	multiples[NeumannDirectionRight] = [2]int{1, 0}
	multiples[NeumannDirectionDown] = [2]int{0, 1}
	multiples[NeumannDirectionLeft] = [2]int{-1, 0}
	return multiples
}
