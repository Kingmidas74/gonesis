package world

import (
	"errors"
	"fmt"
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
)

type World struct {
	contracts.Terrain
	commands []contracts.Command
}

func New(terrain contracts.Terrain, commands []contracts.Command) *World {
	return &World{
		Terrain:  terrain,
		commands: commands,
	}
}

func (w *World) Agents() []contracts.Agent {
	agents := make([]contracts.Agent, 0, len(w.Cells()))
	for _, cell := range w.Cells() {
		if cell.IsAgent() {
			agents = append(agents, cell.Agent())
		}
	}
	return agents
}

func (w *World) Width() int {
	return w.Terrain.Width()
}

func (w *World) Height() int {
	return w.Terrain.Height()
}

func (w *World) Command(commandIdentifier int) contracts.Command {
	if commandIdentifier < 0 || commandIdentifier >= len(w.commands) {
		return nil
	}
	return w.commands[commandIdentifier]
}

func (w *World) Next(config *configuration.Configuration) error {
	return w.runDay(config)
}

func (w *World) runDay(config *configuration.Configuration) error {
	for _, cell := range w.Cells() {
		if !cell.IsAgent() {
			continue
		}
		agent := cell.Agent()
		if err := agent.NextDay(w, w.Command, &config.AgentConfiguration); err != nil {
			return err
		}
	}

	w.removeDeathAgents()
	w.genesis(&config.AgentConfiguration)

	livingAgentsCount := 0
	for _, cell := range w.Cells() {
		if cell.IsAgent() {
			livingAgentsCount++
		}
	}
	if livingAgentsCount > len(w.Cells()) {
		fmt.Println(livingAgentsCount, w.Width()*w.Height())
		return errors.New("too many agents")
	}
	return nil
}

func (w *World) removeDeathAgents() {
	for _, cell := range w.Cells() {
		if !cell.IsAgent() {
			continue
		}
		agent := cell.Agent()
		if !agent.IsAlive() {
			cell.RemoveAgent()
		}
	}
}

func (w *World) genesis(config *configuration.AgentConfiguration) {
	for _, cell := range w.Cells() {
		if !cell.IsAgent() {
			continue
		}
		agent := cell.Agent()
		if !agent.IsAlive() {
			continue
		}

		if children := agent.CreateChildren(w, config); children != nil {
			for _, child := range children {
				childCell := w.Cell(child.X(), child.Y())
				if childCell.IsEmpty() {
					childCell.SetAgent(child)
				}
			}
		}
	}
}
