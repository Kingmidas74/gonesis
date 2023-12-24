package eat

import (
	contract "github.com/kingmidas74/gonesis-engine/internal/domain/contract"
)

func (c *Command) IsInterrupt() bool {
	return c.isInterrupt
}

func (c *Command) Handle(agent contract.Agent, terra contract.Terrain) (int, error) {
	//return 1, nil
	if agent.Type() == contract.AgentTypePlant {
		panic("Plant can't eat")
	}

	availableFood := agent.AvailableFood()

	whereAddress := agent.Address() + 1
	direction := agent.Command(&whereAddress)
	targetCell := terra.GetNeighbor(agent.X(), agent.Y(), direction)
	if targetCell == nil {
		return 1, nil
	}

	targetAgent := targetCell.Agent()
	if targetAgent == nil {
		return 1, nil
	}

	for foodType, energyCost := range availableFood {
		currentCell := terra.Cell(agent.X(), agent.Y())
		if targetCell.Agent().IsAlive() && terra.CanMoveTo(currentCell, targetCell) && targetCell.Agent().Type() == foodType {
			targetEnergy := targetCell.Agent().Energy()
			if targetEnergy < energyCost {
				energyCost = targetEnergy
			}
			agent.IncreaseEnergy(energyCost)
			targetCell.Agent().DecreaseEnergy(energyCost)
			if !targetCell.Agent().IsAlive() {
				targetCell.RemoveAgent()
				targetCell.SetAgent(agent)
				currentCell.RemoveAgent()
			}
			break
		}
	}

	return 1, nil
	/*localDelta := targetCell.CellType().Value() + 1
	deltaAddress := agent.Address() + localDelta
	delta := agent.Command(&deltaAddress)
	return delta, nil*/
}
