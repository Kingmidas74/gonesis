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
	result := make([]contracts.Coords, 0)

	for _, coords := range t.getCoordsMultiples() {
		result = append(result, entity.NewCoords(x+coords[0], y+coords[1]))
	}

	return result
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
