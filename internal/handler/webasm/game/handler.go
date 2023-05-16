//go:build js && wasm

package game

import (
	"syscall/js"

	"github.com/kingmidas74/gonesis-engine/internal/handler/webasm/model"
)

func (h *Handler) InitWorld(args []js.Value) (*model.World, error) {
	width := args[0].Int()
	height := args[1].Int()
	agentsCount := args[2].Int()

	world, err := h.gameService.InitWorld(width, height, agentsCount)
	if err != nil {
		return nil, err
	}

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

	return &model.World{
		Width:  world.Width(),
		Height: world.Height(),
		Cells:  cells,
		Agents: agents,
	}, nil
}
