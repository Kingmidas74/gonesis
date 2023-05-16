package mapper

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/handler/webasm/model"
)

func NewWorld(world contracts.World) model.World {
	cells := make([]model.Cell, len(world.Cells()))
	for i, cell := range world.Cells() {
		cells[i] = model.Cell{
			CellType: cell.CellType().Value(),
		}
	}

	agents := make([]model.Agent, len(world.Agents()))
	for i, agent := range world.Agents() {
		agents[i] = model.Agent{
			X:      agent.X(),
			Y:      agent.Y(),
			Energy: agent.Energy(),
		}
	}

	return model.World{
		Width:  world.Width(),
		Height: world.Height(),
		Cells:  cells,
		Agents: agents,
	}
}
