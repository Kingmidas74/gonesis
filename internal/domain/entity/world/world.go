package world

import (
	"fmt"
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
)

type World struct {
	contracts.Terrain
	commands   []contracts.Command
	currentDay int
}

func New(terrain contracts.Terrain, commands []contracts.Command) *World {
	return &World{
		Terrain:    terrain,
		commands:   commands,
		currentDay: 0,
	}
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
	if err := w.runDay(config); err != nil {
		panic(err.Error())
		return err
	}

	w.currentDay++
	return nil
}

func (w *World) CurrentDay() int {
	return w.currentDay
}

func (w *World) runDay(config *configuration.Configuration) error {
	handledAgents := make(map[string]struct{})
	livingAgentsCount := 0

	for _, cell := range w.Cells() {
		if !cell.IsAgent() {
			continue
		}

		agent := cell.Agent()
		if _, ok := handledAgents[agent.ID()]; ok {
			continue
		}

		if err := agent.NextDay(w, w.Command); err != nil {
			return err
		}

		if !agent.IsAlive() {
			cell.RemoveAgent()
			continue
		}

		children := agent.CreateChildren(w, config)
		for _, child := range children {
			handledAgents[child.ID()] = struct{}{}
		}
		handledAgents[agent.ID()] = struct{}{}

		livingAgentsCount++
	}

	if livingAgentsCount > len(w.Cells()) {
		return fmt.Errorf("too many agents: current %v, max %v", livingAgentsCount, w.Width()*w.Height())
	}

	return nil

	/*livingAgentsCount := 0
	for _, cell := range w.Cells() {
		if !cell.IsAgent() {
			continue
		}
		agent := cell.Agent()
		if err := agent.NextDay(w, w.Command, config); err != nil {
			return err
		}
		if !agent.IsAlive() {
			cell.RemoveAgent()
			continue
		}
		livingAgentsCount++
		if children := agent.CreateChildren(w, config); children != nil {
			for _, child := range children {
				childCell := w.Cell(child.X(), child.Y())
				if childCell.IsEmpty() {
					livingAgentsCount++
					childCell.SetAgent(child)
				}
			}
		}
	}

	if livingAgentsCount > len(w.Cells()) {
		fmt.Println(livingAgentsCount, w.Width()*w.Height())
		panic("too many agents")
		return errors.New("too many agents")
	}
	return nil

	*/
}
