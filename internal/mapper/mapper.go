package mapper

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/contract"
	"github.com/kingmidas74/gonesis-engine/internal/mapper/model"
)

func NewWorld(world contract.World) model.World {
	cells := make([]model.Cell, len(world.Cells()))
	for i, cell := range world.Cells() {
		cells[i] = model.Cell{
			X:         cell.X(),
			Y:         cell.Y(),
			NorthWall: cell.NorthWall(),
			SouthWall: cell.SouthWall(),
			WestWall:  cell.WestWall(),
			EastWall:  cell.EastWall(),
		}

		if cell.Agent() != nil {
			cells[i].Agent = &model.Agent{
				X:          cell.Agent().X(),
				Y:          cell.Agent().Y(),
				Energy:     cell.Agent().Energy(),
				AgentType:  cell.Agent().Type().String(),
				Generation: cell.Agent().Generation(),
				Brain: model.Brain{
					Commands: cell.Agent().Brain().Commands(),
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
