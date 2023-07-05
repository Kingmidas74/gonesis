package contracts

import "github.com/kingmidas74/gonesis-engine/internal/domain/enum"

type Cell interface {
	Energy
	Coords

	X() int
	Y() int

	CellType() enum.CellType
	SetCellType(cellType enum.CellType)

	Agent() Agent
	SetAgent(a Agent)
	RemoveAgent()

	IsEmpty() bool
	IsAgent() bool

	NorthWall() bool
	SetNorthWall(flag bool)

	SouthWall() bool
	SetSouthWall(flag bool)

	WestWall() bool
	SetWestWall(flag bool)

	EastWall() bool
	SetEastWall(flag bool)
}
