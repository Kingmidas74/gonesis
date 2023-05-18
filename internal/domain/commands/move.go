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
	whereAddress := agent.Address() + 1
	direction := agent.Command(&whereAddress)
	neighborCell := terra.GetNeighbor(agent.X(), agent.Y(), direction)
	if neighborCell == nil {
		return 0
	}
	if neighborCell.CellType() == enum.CellTypeEmpty {
		agent.SetX(neighborCell.X())
		agent.SetY(neighborCell.Y())
	}
	localDelta := neighborCell.CellType().Value() + 1
	deltaAddress := agent.Address() + localDelta
	delta := agent.Command(&deltaAddress)
	return delta
}

func (c *MoveCommand) IsInterrupt() bool {
	return c.isInterrupt
}
