package commands

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
)

type EndSubroutineCommand struct {
	isInterrupt bool
}

func NewEndSubroutineCommand() *EndSubroutineCommand {
	return &EndSubroutineCommand{
		isInterrupt: false,
	}
}

func (c *EndSubroutineCommand) Handle(agent contracts.Agent, terra contracts.Terrain) int {
	agent.Return()
	return 1
}

func (c *EndSubroutineCommand) IsInterrupt() bool {
	return c.isInterrupt
}
