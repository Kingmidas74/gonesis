package world

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/terrain"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type World struct {
	terrain.Terrain
}

func (w *World) Action(maxSteps int, callback func(terrain.Terrain, int)) error {
	callback(w.Terrain, 0)
	for currentDay, livingAgents := 1, w.filterLivingAgents(); len(livingAgents) > 0; currentDay, livingAgents = currentDay+1, w.filterLivingAgents() {
		err := w.runDay(livingAgents, maxSteps)
		if err != nil {
			return err
		}
		w.cleanDeath()
		callback(w.Terrain, currentDay)
	}
	return nil
}

func (w *World) runDay(agents []*agent.Agent, maxSteps int) error {
	for _, a := range agents {
		err := a.NextDay(maxSteps, func(i int) agent.Command {
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *World) filterLivingAgents() []*agent.Agent {
	result := make([]*agent.Agent, 0)
	for _, cell := range w.GetCells() {
		if cell.Agent() != nil && cell.Agent().IsAlive() {
			result = append(result, cell.Agent())
		}
	}
	return result
}

func (w *World) cleanDeath() {
	for y := 0; y < w.GetHeight(); y++ {
		for x := 0; x < w.GetWidth(); x++ {
			currentCell := w.GetCell(x, y)
			if currentCell.Agent() != nil && !currentCell.Agent().IsAlive() {
				currentCell.SetAgent(nil)
				currentCell.SetCellType(enum.CellTypeEmpty)
			}
		}
	}
}
