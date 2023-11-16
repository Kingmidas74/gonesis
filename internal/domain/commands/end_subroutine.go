package commands

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type EndSubroutineCommand struct {
	isInterrupt bool
	available   []enum.AgentType
}

func NewEndSubroutineCommand() *EndSubroutineCommand {
	return &EndSubroutineCommand{
		isInterrupt: false,
		available:   []enum.AgentType{enum.AgentTypePlant, enum.AgentTypeHerbivore, enum.AgentTypeCarnivore, enum.AgentTypeOmnivore},
	}
}

func (c *EndSubroutineCommand) Handle(agent contracts.Agent, terra contracts.Terrain) int {
	agent.Return()
	return 1
}

func (c *EndSubroutineCommand) IsInterrupt() bool {
	return c.isInterrupt
}

func (c *EndSubroutineCommand) IsAvailable(agent contracts.AgentNature) bool {
	for i := range c.available {
		if c.available[i] == agent.AgentType() {
			return true
		}
	}
	return false
}
