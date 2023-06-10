package entity

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"sync"
)

type Coords struct {
	x int
	y int

	mu sync.Mutex
}

func NewCoords(x, y int) contracts.Coords {
	return &Coords{
		x: x,
		y: y,
	}
}

func (c *Coords) X() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.x
}

func (c *Coords) SetX(x int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.x = x
}

func (c *Coords) Y() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.y
}

func (c *Coords) SetY(y int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.y = y
}
