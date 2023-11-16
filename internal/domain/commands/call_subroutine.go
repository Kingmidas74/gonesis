package commands

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type CallSubroutineCommand struct {
	isInterrupt bool
	available   []enum.AgentType
}

func NewCallSubroutineCommand() *CallSubroutineCommand {
	return &CallSubroutineCommand{
		isInterrupt: false,
		available:   []enum.AgentType{enum.AgentTypePlant, enum.AgentTypeHerbivore, enum.AgentTypeCarnivore, enum.AgentTypeOmnivore},
	}
}

func (c *CallSubroutineCommand) Handle(agent contracts.Agent, terra contracts.Terrain) int {
	fromAddress, toAddress := agent.Address()+1, agent.Address()+2
	fromCommand, lengthCommands := agent.Command(&fromAddress), agent.Command(&toAddress)
	if lengthCommands <= 0 {
		lengthCommands = len(agent.Commands())
	}

	agent.KeepAddress(fromCommand, lengthCommands)
	return 1
}

func (c *CallSubroutineCommand) IsInterrupt() bool {
	return c.isInterrupt
}

func (c *CallSubroutineCommand) IsAvailable(agent contracts.AgentNature) bool {
	for i := range c.available {
		if c.available[i] == agent.AgentType() {
			return true
		}
	}
	return false
}
