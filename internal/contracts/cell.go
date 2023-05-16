package contracts

import "github.com/kingmidas74/gonesis-engine/internal/domain/enum"

type Cell interface {
	X() int
	Y() int
	CellType() enum.CellType
	SetCellType(cellType enum.CellType)
}
