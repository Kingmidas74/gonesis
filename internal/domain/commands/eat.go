package commands

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type EatCommand struct {
	isInterrupt bool
	available   []enum.AgentType
}

func NewEatCommand() *EatCommand {
	return &EatCommand{
		isInterrupt: true,
		available:   []enum.AgentType{enum.AgentTypeHerbivore, enum.AgentTypeCarnivore, enum.AgentTypeOmnivore},
	}
}

func (c *EatCommand) Handle(agent contracts.Agent, terra contracts.Terrain) int {
	if agent.AgentType() == enum.AgentTypePlant {
		panic("Plant can't eat")
	}

	availableFood := agent.AvailableFood()

	whereAddress := agent.Address() + 1
	direction := agent.Command(&whereAddress)
	targetCell := terra.GetNeighbor(agent.X(), agent.Y(), direction)
	if targetCell == nil {
		return 1
	}

	if !targetCell.IsAgent() {
		return 1
	}

	for foodType, energyCost := range availableFood {
		if targetCell.Agent().AgentType() == foodType && targetCell.Agent().IsAlive() {
			targetEnergy := targetCell.Agent().Energy()
			if targetEnergy < energyCost {
				energyCost = targetEnergy
			}

			agent.IncreaseEnergy(energyCost)
			targetCell.Agent().DecreaseEnergy(energyCost)

			if !targetCell.Agent().IsAlive() {
				targetCell.Agent().Kill(terra)
				terra.Cell(agent.X(), agent.Y()).RemoveAgent()
				targetCell.SetAgent(agent)
			}

			break
		}
	}

	localDelta := enum.CellTypeAgent.Value() + 1
	deltaAddress := agent.Address() + localDelta
	delta := agent.Command(&deltaAddress)
	return delta
}

func (c *EatCommand) IsInterrupt() bool {
	return c.isInterrupt
}

func (c *EatCommand) IsAvailable(agent contracts.AgentNature) bool {
	for i := range c.available {
		if c.available[i] == agent.AgentType() {
			return true
		}
	}
	return false
}
