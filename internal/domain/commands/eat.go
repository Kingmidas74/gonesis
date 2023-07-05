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
	if agent.AgentType() == enum.AgentTypePlant {
		agent.IncreaseEnergy(terra.Cell(agent.X(), agent.Y()).Energy() / 2)
	} else {
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
		terra.Cell(agent.X(), agent.Y()).RemoveAgent()
		targetCell.SetAgent(agent)
	}
	localDelta := enum.CellTypeAgent.Value() + 1
	deltaAddress := agent.Address() + localDelta
	delta := agent.Command(&deltaAddress)
	return delta
}

func (c *EatCommand) IsInterrupt() bool {
	return c.isInterrupt
}
