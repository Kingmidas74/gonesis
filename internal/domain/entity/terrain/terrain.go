package terrain

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"github.com/kingmidas74/gonesis-engine/internal/util"
)

type Terrain struct {
	Cells  []entity.Cell
	Width  int
	Height int
}

func (t *Terrain) transformX(x int) int {
	return util.ModLikePython(x, t.Width)
}

func (t *Terrain) transformY(y int) int {
	return util.ModLikePython(y, t.Height)
}

func (t *Terrain) getCellIndex(x, y int) int {
	x = t.transformX(x)
	y = t.transformY(y)
	return y*t.Width + x
}

func (t *Terrain) GetCell(x, y int) entity.Cell {
	return t.Cells[t.getCellIndex(x, y)]
}

func (t *Terrain) GetCells() []entity.Cell {
	return t.Cells
}

func (t *Terrain) GetWidth() int {
	return t.Width
}

func (t *Terrain) GetHeight() int {
	return t.Height
}

func (t *Terrain) GetNeighbor(x, y int, direction int) entity.Cell {
	multiples := t.getCoordsMultiples()
	neumannDirection := enum.Direction(util.ModLikePython(direction, len(multiples)))
	return t.GetCell(x+multiples[neumannDirection][0], y+multiples[neumannDirection][1])
}

func (t *Terrain) GetNeighbors(x, y int) []entity.Cell {
	result := make([]entity.Cell, 0)

	for _, coords := range t.getCoordsMultiples() {
		result = append(result, t.GetCell(x+coords[0], y+coords[1]))
	}
	return result
}

func (t *Terrain) getCoordsMultiples() map[enum.Direction][2]int {
	return map[enum.Direction][2]int{
		enum.DirectionUp:    {0, -1},
		enum.DirectionRight: {1, 0},
		enum.DirectionDown:  {0, 1},
		enum.DirectionLeft:  {-1, 0},
	}
}
