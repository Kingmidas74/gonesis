package entity

import "github.com/kingmidas74/gonesis-engine/internal/domain/enum"

type Cell struct {
	Coords

	cellType enum.CellType
}

func (c *Cell) CellType() enum.CellType {
	return c.cellType
}

func (c *Cell) SetCellType(cellType enum.CellType) {
	c.cellType = cellType
}
