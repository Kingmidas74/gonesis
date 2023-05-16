package commands

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type MoveCommand struct {
	isInterrupt bool
}

func NewMoveCommand() *MoveCommand {
	return &MoveCommand{}
}

func (c *MoveCommand) Handle(agent contracts.Agent, terra contracts.Terrain) int {
	whereAddress := agent.Address() + 1
	direction := agent.Command(&whereAddress)
	neighborCell := terra.GetNeighbor(agent.X(), agent.Y(), direction)
	switch neighborCell.CellType() {
	case enum.CellTypeEmpty:
		agent.SetX(neighborCell.X())
		agent.SetY(neighborCell.Y())
	}
	return 1
}

func (c *MoveCommand) IsInterrupt() bool {
	return c.isInterrupt
}
