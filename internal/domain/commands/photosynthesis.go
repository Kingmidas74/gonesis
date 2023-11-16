package commands

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type PhotosynthesisCommand struct {
	isInterrupt bool
	available   []enum.AgentType
}

func NewPhotosynthesisCommand() *PhotosynthesisCommand {
	return &PhotosynthesisCommand{
		isInterrupt: true,
		available:   []enum.AgentType{enum.AgentTypePlant},
	}
}

func (c *PhotosynthesisCommand) Handle(agent contracts.Agent, terra contracts.Terrain) int {
	agent.IncreaseEnergy(1)
	return 1
}

func (c *PhotosynthesisCommand) IsInterrupt() bool {
	return c.isInterrupt
}

func (c *PhotosynthesisCommand) IsAvailable(agent contracts.AgentNature) bool {
	for i := range c.available {
		if c.available[i] == agent.AgentType() {
			return true
		}
	}
	return false
}
