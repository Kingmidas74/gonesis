package commands

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
)

type Photosynthesis struct {
	isInterrupt bool
}

func NewPhotosynthesis() *Photosynthesis {
	return &Photosynthesis{}
}

func (c *Photosynthesis) Handle(agent contracts.Agent, terra contracts.Terrain) int {
	return 1
}

func (c *Photosynthesis) IsInterrupt() bool {
	return c.isInterrupt
}
