package end_subroutine

import (
	contract "github.com/kingmidas74/gonesis-engine/internal/domain/contract"
)

type EndSubroutineCommand struct {
	isInterrupt bool
}

func NewEndSubroutineCommand() *EndSubroutineCommand {
	return &EndSubroutineCommand{
		isInterrupt: false,
	}
}

func (c *EndSubroutineCommand) Handle(agent contract.Agent, terra contract.Terrain) int {
	//agent.Return()
	return 1
}

func (c *EndSubroutineCommand) IsInterrupt() bool {
	return c.isInterrupt
}
