package commands

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
)

type PhotosynthesisCommand struct {
	isInterrupt bool
}

func NewPhotosynthesisCommand() *PhotosynthesisCommand {
	return &PhotosynthesisCommand{
		isInterrupt: true,
	}
}

func (c *PhotosynthesisCommand) Handle(agent contracts.Agent, terra contracts.Terrain) int {
	agent.IncreaseEnergy(1)
	return 1
}

func (c *PhotosynthesisCommand) IsInterrupt() bool {
	return c.isInterrupt
}
