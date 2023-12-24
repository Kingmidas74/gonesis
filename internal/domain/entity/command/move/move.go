package move

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/contract"
)

func (c *Command) IsInterrupt() bool {
	return c.isInterrupt
}

func (c *Command) Handle(agent contract.Agent, terra contract.Terrain) (int, error) {
	whereAddress := agent.Address() + 1
	direction := agent.Command(&whereAddress)
	targetCell := terra.GetNeighbor(agent.X(), agent.Y(), direction)
	if targetCell == nil {
		panic("target cell is nil")
	}
	currentCell := terra.Cell(agent.X(), agent.Y())

	//originalTargetCellType := targetCell.CellType()
	if targetCell.IsEmpty() && terra.CanMoveTo(currentCell, targetCell) {
		currentCell.RemoveAgent()
		targetCell.SetAgent(agent)
	}
	return 1, nil
	/*
		localDelta := originalTargetCellType.Value() + 1
		deltaAddress := agent.Address() + localDelta
		delta := agent.Command(&deltaAddress)

		return delta, nil
	*/
}
