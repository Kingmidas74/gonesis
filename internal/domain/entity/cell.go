package entity

import "github.com/kingmidas74/gonesis-engine/internal/domain/enum"

type Cell struct {
	Coords

	CellType enum.CellType
}

func (this *Cell) GetCellType() enum.CellType {
	return this.CellType
}

func (this *Cell) SetCellType(cellType enum.CellType) {
	this.CellType = cellType
}
