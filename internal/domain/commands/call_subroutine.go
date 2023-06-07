package commands

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
)

type CallSubroutineCommand struct {
	isInterrupt bool
}

func NewCallSubroutineCommand() *CallSubroutineCommand {
	return &CallSubroutineCommand{
		isInterrupt: false,
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
