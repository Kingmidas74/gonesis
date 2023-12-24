package call_subroutine

import (
	contract "github.com/kingmidas74/gonesis-engine/internal/domain/contract"
)

func (c *Command) Handle(agent contract.Agent, terra contract.Terrain) (int, error) {
	/*fromAddress, toAddress := agent.Address()+1, agent.Address()+2
	fromCommand, lengthCommands := agent.Command(&fromAddress), agent.Command(&toAddress)
	if lengthCommands <= 0 {
		lengthCommands = len(agent.Commands())
	}

	agent.KeepAddress(fromCommand, lengthCommands)*/
	return 1, nil
}

func (c *Command) IsInterrupt() bool {
	return c.isInterrupt
}
