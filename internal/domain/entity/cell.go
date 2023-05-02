package entity

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Cell struct {
	Coords

	cellType enum.CellType
	agent    *agent.Agent
}

func (c *Cell) CellType() enum.CellType {
	return c.cellType
}

func (c *Cell) SetCellType(cellType enum.CellType) {
	c.cellType = cellType
}

func (c *Cell) Agent() *agent.Agent {
	return c.agent
}

func (c *Cell) SetAgent(a *agent.Agent) {
	c.agent = a
}
