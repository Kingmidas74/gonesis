package commands

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type MoveCommand struct {
	isInterrupt bool
}

func NewMoveCommand() *MoveCommand {
	return &MoveCommand{
		isInterrupt: true,
	}
}

func (c *MoveCommand) Handle(agent contracts.Agent, terra contracts.Terrain) int {
	switch agent.AgentType() {
	case enum.AgentTypeHerbivore:
		return c.handleHerbivore(agent, terra)
	case enum.AgentTypeCarnivore:
		return c.handleCarnivore(agent, terra)
	case enum.AgentTypePlant:
		return c.handlePlant(agent, terra)
	case enum.AgentTypeOmnivore:
		return c.handleHerbivore(agent, terra)
	}
	return 1
}

func (c *MoveCommand) IsInterrupt() bool {
	return c.isInterrupt
}

func (c *MoveCommand) handleHerbivore(agent contracts.Agent, terra contracts.Terrain) int {
	whereAddress := agent.Address() + 1
	direction := agent.Command(&whereAddress)
	targetCell := terra.GetNeighbor(agent.X(), agent.Y(), direction)
	if targetCell == nil {
		return 0
	}

	originalTargetCellType := targetCell.CellType()
	if targetCell.IsEmpty() {
		terra.Cell(agent.X(), agent.Y()).RemoveAgent()
		targetCell.SetAgent(agent)
	}

	localDelta := originalTargetCellType.Value() + 1
	deltaAddress := agent.Address() + localDelta
	delta := agent.Command(&deltaAddress)
	return delta
}

func (c *MoveCommand) handleCarnivore(agent contracts.Agent, terra contracts.Terrain) int {
	whereAddress := agent.Address() + 1
	direction := agent.Command(&whereAddress)
	targetCell := terra.GetNeighbor(agent.X(), agent.Y(), direction)
	if targetCell == nil {
		return 0
	}

	originalTargetCellType := targetCell.CellType()
	if targetCell.IsEmpty() {
		terra.Cell(agent.X(), agent.Y()).RemoveAgent()
		targetCell.SetAgent(agent)
	}

	localDelta := originalTargetCellType.Value() + 1
	deltaAddress := agent.Address() + localDelta
	delta := agent.Command(&deltaAddress)
	return delta
}

func (c *MoveCommand) handlePlant(agent contracts.Agent, terra contracts.Terrain) int {
	return 1
}
