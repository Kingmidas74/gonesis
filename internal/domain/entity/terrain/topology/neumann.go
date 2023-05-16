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
	neumannDirection := NeumannDirection(util.ModLikePython(direction, len(t.getCoordsMultiples())))

	return entity.NewCoords(x+multiples[neumannDirection][0], y+multiples[neumannDirection][1])
}

func (t NeumannTopology) GetNeighbors(x, y int) []contracts.Coords {
	result := make([]contracts.Coords, 0)

	for _, coords := range t.getCoordsMultiples() {
		result = append(result, entity.NewCoords(x+coords[0], y+coords[1]))
	}

	return result
}

func (t NeumannTopology) getCoordsMultiples() map[NeumannDirection][2]int {
	multiples := make(map[NeumannDirection][2]int)
	multiples[NeumannDirectionUp] = [2]int{0, -1}
	multiples[NeumannDirectionRight] = [2]int{1, 0}
	multiples[NeumannDirectionDown] = [2]int{0, 1}
	multiples[NeumannDirectionLeft] = [2]int{-1, 0}
	return multiples
}
