package commands

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type EatCommand struct {
	isInterrupt bool
}

func NewEatCommand() *EatCommand {
	return &EatCommand{
		isInterrupt: true,
	}
}

func (c *EatCommand) Handle(agent contracts.Agent, terra contracts.Terrain) int {
	whereAddress := agent.Address() + 1
	direction := agent.Command(&whereAddress)
	targetCell := terra.GetNeighbor(agent.X(), agent.Y(), direction)
	if targetCell == nil {
		return 1
	}

	if !targetCell.IsAgent() {
		return 1
	}

	agent.IncreaseEnergy(targetCell.Agent().Energy())
	targetCell.Agent().Kill(terra)
	targetCell.SetAgent(agent)

	localDelta := enum.CellTypeAgent.Value() + 1
	deltaAddress := agent.Address() + localDelta
	delta := agent.Command(&deltaAddress)
	return delta

	return 1
}

func (c *EatCommand) IsInterrupt() bool {
	return c.isInterrupt
}
