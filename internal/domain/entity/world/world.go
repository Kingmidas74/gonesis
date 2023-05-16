package world

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type World struct {
	contracts.Terrain

	agents   []contracts.Agent
	commands []contracts.Command
}

func New(terrain contracts.Terrain, agents []contracts.Agent, commands []contracts.Command) contracts.World {
	return &World{
		Terrain:  terrain,
		agents:   agents,
		commands: commands,
	}
}

func (w *World) Action(maxSteps int, callback func(contracts.World, int)) error {
	callback(w, 0)
	w.updateCells()
	for currentDay, livingAgentsCount := 1, w.countLivingAgents(); livingAgentsCount > 0; currentDay, livingAgentsCount = currentDay+1, w.countLivingAgents() {
		err := w.runDay(maxSteps)
		if err != nil {
			return err
		}
		w.updateCells()
		callback(w, currentDay)
	}
	return nil
}

func (w *World) Agents() []contracts.Agent {
	return w.agents
}

func (w *World) Width() int {
	return w.Terrain.Width()
}

func (w *World) Height() int {
	return w.Terrain.Height()
}

func (w *World) Command(commandIdentifier int) contracts.Command {
	return w.commands[commandIdentifier]
}

func (w *World) runDay(maxSteps int) error {
	for _, a := range w.Agents() {
		if err := a.NextDay(maxSteps, w, w.Command); err != nil {
			return err
		}
	}
	return nil
}

func (w *World) countLivingAgents() int {
	result := 0
	for _, a := range w.Agents() {
		if a.IsAlive() {
			result++
		}
	}
	return result
}

func (w *World) updateCells() {
	for _, c := range w.Cells() {
		if c.CellType() != enum.CellTypeObstacle {
			c.SetCellType(enum.CellTypeEmpty)
		}
	}

	for _, a := range w.Agents() {
		if a.IsAlive() {
			w.SetCellType(a.X(), a.Y(), enum.CellTypeOrganic)
		}
	}
}
