package commands

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type MoveCommand struct {
	isInterrupt bool
	available   []enum.AgentType
}

func NewMoveCommand() *MoveCommand {
	return &MoveCommand{
		isInterrupt: true,
		available:   []enum.AgentType{enum.AgentTypeHerbivore, enum.AgentTypeCarnivore, enum.AgentTypeOmnivore},
	}
}

func (c *MoveCommand) IsInterrupt() bool {
	return c.isInterrupt
}

func (c *MoveCommand) Handle(agent contracts.Agent, terra contracts.Terrain) int {
	whereAddress := agent.Address() + 1
	direction := agent.Command(&whereAddress)
	targetCell := terra.GetNeighbor(agent.X(), agent.Y(), direction)
	if targetCell == nil {
		panic("target cell is nil")
	}

	originalTargetCellType := targetCell.CellType()
	if terra.CanMoveTo(terra.Cell(agent.X(), agent.Y()), targetCell) {
		terra.Cell(agent.X(), agent.Y()).RemoveAgent()
		targetCell.SetAgent(agent)
	}

	localDelta := originalTargetCellType.Value() + 1
	deltaAddress := agent.Address() + localDelta
	delta := agent.Command(&deltaAddress)
	return delta
}

func (c *MoveCommand) IsAvailable(agent contracts.AgentNature) bool {
	for i := range c.available {
		if c.available[i] == agent.AgentType() {
			return true
		}
	}
	return false
}
