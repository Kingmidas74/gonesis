package photosynthesis

import (
	contract "github.com/kingmidas74/gonesis-engine/internal/domain/contract"
)

func (c *Command) IsInterrupt() bool {
	return c.isInterrupt
}

func (c *Command) Handle(agent contract.Agent, terra contract.Terrain) (int, error) {
	if agent.Type() != contract.AgentTypePlant {
		panic("Only plant can photosynthesis")
	}

	agent.IncreaseEnergy(1)
	return 1, nil
}
