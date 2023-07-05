package mapper

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/mapper/model"
)

func NewWorld(world contracts.World) model.World {
	cells := make([]model.Cell, len(world.Cells()))
	for i, cell := range world.Cells() {
		cells[i] = model.Cell{
			CellType:  cell.CellType().String(),
			Energy:    cell.Energy(),
			X:         cell.X(),
			Y:         cell.Y(),
			NorthWall: cell.NorthWall(),
			SouthWall: cell.SouthWall(),
			WestWall:  cell.WestWall(),
			EastWall:  cell.EastWall(),
		}

		if cell.IsAgent() {
			cells[i].Agent = &model.Agent{
				X:          cell.Agent().X(),
				Y:          cell.Agent().Y(),
				Energy:     cell.Agent().Energy(),
				AgentType:  cell.Agent().AgentType().String(),
				Generation: cell.Agent().Generation(),
				Brain: model.Brain{
					Commands: cell.Agent().Commands(),
				},
			}
		}
	}

	return model.World{
		Width:      world.Width(),
		Height:     world.Height(),
		Cells:      cells,
		CurrentDay: world.CurrentDay(),
	}
}

func NewAgent(a contracts.Agent) model.Agent {
	return model.Agent{
		X:         a.X(),
		Y:         a.Y(),
		Energy:    a.Energy(),
		AgentType: a.AgentType().String(),

		Brain: model.Brain{
			Commands: a.Commands(),
		},
	}
}
