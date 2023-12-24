package cell

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/contract"
)

func (c *Cell) SetAgent(a contract.Agent) {
	if c.agent != nil && c.agent.IsAlive() {
		panic("cell already has a live agent")
	}

	if c.agent != nil {
		panic("cell already has a dead agent")
	}

	c.agent = a
	a.SetX(c.x)
	a.SetY(c.y)
}

func (c *Cell) RemoveAgent() {
	c.agent = nil
}

func (c *Cell) IsEmpty() bool {
	return c.agent == nil
}

func (c *Cell) Agent() contract.Agent {
	return c.agent
}

func (c *Cell) NorthWall() bool {
	return c.northWall
}

func (c *Cell) SouthWall() bool {
	return c.southWall
}

func (c *Cell) WestWall() bool {
	return c.westWall
}

func (c *Cell) EastWall() bool {
	return c.eastWall
}

func (c *Cell) X() int {
	return c.x
}

func (c *Cell) Y() int {
	return c.y
}
