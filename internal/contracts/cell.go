package contracts

import "github.com/kingmidas74/gonesis-engine/internal/domain/enum"

type Cell interface {
	Energy

	X() int
	Y() int
	CellType() enum.CellType
	SetCellType(cellType enum.CellType)

	Agent() Agent
	SetAgent(a Agent)
	RemoveAgent()

	IsEmpty() bool
	IsAgent() bool

	Lock()
	Unlock()
}
