package entity

import (
	"sync"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Cell struct {
	Coords

	mu sync.Mutex

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
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cellType
}

func (c *Cell) SetCellType(cellType enum.CellType) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cellType = cellType
}

func (c *Cell) Energy() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.energy
}

func (c *Cell) IncreaseEnergy(delta int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.energy += delta
}

func (c *Cell) DecreaseEnergy(delta int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.energy -= delta
}

func (c *Cell) Agent() contracts.Agent {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.agent
}

func (c *Cell) SetAgent(a contracts.Agent) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cellType = enum.CellTypeAgent
	c.agent = a
	a.SetX(c.X())
	a.SetY(c.Y())
}

func (c *Cell) RemoveAgent() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cellType = enum.CellTypeEmpty
	c.agent = nil
}

func (c *Cell) IsEmpty() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cellType == enum.CellTypeEmpty
}
func (c *Cell) IsAgent() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cellType == enum.CellTypeAgent
}

func (c *Cell) Lock() {
	c.mu.Lock()
}

func (c *Cell) Unlock() {
	c.mu.Unlock()
}
