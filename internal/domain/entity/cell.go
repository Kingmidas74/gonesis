package entity

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Cell struct {
	Coords

	cellType enum.CellType
	energy   int

	agent contracts.Agent
}

func NewCell(x, y int, cellType enum.CellType) *Cell {
	return &Cell{
		Coords:   Coords{x: x, y: y},
		cellType: cellType,
		energy:   0,
	}
}

func (c *Cell) CellType() enum.CellType {
	return c.cellType
}

func (c *Cell) SetCellType(cellType enum.CellType) {
	c.cellType = cellType
}

func (c *Cell) Energy() int {
	return c.energy
}

func (c *Cell) IncreaseEnergy(delta int) {
	c.energy += delta
}

func (c *Cell) DecreaseEnergy(delta int) {
	c.energy -= delta
}

func (c *Cell) Agent() contracts.Agent {
	return c.agent
}

func (c *Cell) SetAgent(a contracts.Agent) {
	c.cellType = enum.CellTypeAgent
	c.agent = a
	a.SetX(c.X())
	a.SetY(c.Y())
}

func (c *Cell) RemoveAgent() {
	c.cellType = enum.CellTypeEmpty
	c.agent = nil
}

func (c *Cell) IsEmpty() bool {
	return c.cellType == enum.CellTypeEmpty
}
func (c *Cell) IsAgent() bool {
	return c.cellType == enum.CellTypeAgent
}
