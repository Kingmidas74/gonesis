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

	northWall bool
	southWall bool
	westWall  bool
	eastWall  bool
}

func NewCell(x, y int) *Cell {
	return &Cell{
		Coords:   Coords{x: x, y: y},
		cellType: enum.CellTypeEmpty,
		energy:   0,

		northWall: true,
		southWall: true,
		westWall:  true,
		eastWall:  true,
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
	if c.agent != nil {
		panic("cell already has an agent")
	}

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
	return c.cellType == enum.CellTypeAgent && c.agent != nil
}

func (c *Cell) NorthWall() bool {
	return c.northWall
}

func (c *Cell) SetNorthWall(flag bool) {
	c.northWall = flag
}

func (c *Cell) SouthWall() bool {
	return c.southWall
}

func (c *Cell) SetSouthWall(flag bool) {
	c.southWall = flag
}

func (c *Cell) WestWall() bool {
	return c.westWall
}

func (c *Cell) SetWestWall(flag bool) {
	c.westWall = flag
}

func (c *Cell) EastWall() bool {
	return c.eastWall
}

func (c *Cell) SetEastWall(flag bool) {
	c.eastWall = flag
}
