package world

import (
	"fmt"
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"sync"
)

type World struct {
	contracts.Terrain
	commands []contracts.Command

	_agents []contracts.Agent
}

func New(terrain contracts.Terrain, commands []contracts.Command) *World {
	return &World{
		Terrain:  terrain,
		commands: commands,
		_agents:  make([]contracts.Agent, 0, terrain.Width()*terrain.Height()),
	}
}

func (w *World) Agents() []contracts.Agent {
	w._agents = w._agents[:0]
	for _, cell := range w.Cells() {
		if cell.IsAgent() {
			w._agents = append(w._agents, cell.Agent())
		}
	}
	return w._agents
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
	errors := make(chan error)
	done := make(chan bool)

	var wg sync.WaitGroup

	for _, cell := range w.Cells() {
		if !cell.IsAgent() {
			continue
		}
		wg.Add(1)

		go func(cell contracts.Cell) {
			defer wg.Done()

			agent := cell.Agent()

			if err := agent.NextDay(w, w.Command, config); err != nil {
				errors <- err
				return
			}

			if !agent.IsAlive() {
				cell.RemoveAgent()
				return
			}

			_ = agent.CreateChildren(w, config)

		}(cell)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		livingAgentsCount := 0
		for _, cell := range w.Cells() {
			if cell.IsAgent() && cell.Agent().IsAlive() {
				livingAgentsCount++
			}
		}
		if int(livingAgentsCount) > len(w.Cells()) {
			panic(fmt.Errorf("too many agents: current %v, max %v", livingAgentsCount, w.Width()*w.Height()))
			return fmt.Errorf("too many agents: current %v, max %v", livingAgentsCount, w.Width()*w.Height())
		}
		return nil
	case err := <-errors:
		return err
	}
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
