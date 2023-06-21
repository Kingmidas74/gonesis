package agent

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"math/rand"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
)

type Agent struct {
	contracts.Brain
	contracts.AgentNature

	entity.Coords

	energy     int
	generation int

	config *configuration.Configuration
}

func NewAgent(nature contracts.AgentNature) contracts.Agent {
	return &Agent{
		Brain:       NewBrain(nature.BrainVolume()),
		AgentNature: nature,
		energy:      nature.InitialEnergy(),
		generation:  0,
	}
}

func NewAgentWithBrain(nature contracts.AgentNature, brain contracts.Brain, generation int) contracts.Agent {
	return &Agent{
		Brain:       brain,
		AgentNature: nature,
		energy:      nature.InitialEnergy(),
		generation:  generation,
	}
}

func (a *Agent) Energy() int {
	return a.energy
}

func (a *Agent) IsAlive() bool {
	return a.Energy() > 0
}

// TODO: replace findCommandPredicate with []contracts.Command
func (a *Agent) NextDay(terra contracts.Terrain, findCommandPredicate func(int) contracts.Command, config *configuration.Configuration) error {
	for step := 0; a.IsAlive() && step < a.MaxEnergy(); step++ {
		commandIdentifier := a.Command(nil)
		command := findCommandPredicate(commandIdentifier)
		if command == nil {
			a.IncreaseAddress(commandIdentifier)
			continue
		}
		delta := command.Handle(a, terra)
		a.IncreaseAddress(delta)
		if command.IsInterrupt() {
			break
		}
	}
	a.DecreaseEnergy(1)
	return nil
}

func (a *Agent) IncreaseEnergy(delta int) {
	a.energy += delta
	if a.energy > a.MaxEnergy() {
		a.energy = a.MaxEnergy()
	}
}

func (a *Agent) DecreaseEnergy(delta int) {
	a.energy -= delta
}

func (a *Agent) Generation() int {
	return a.generation
}

func (a *Agent) CreateChildren(terra contracts.Terrain, config *configuration.Configuration) []contracts.Agent {
	neighbors := terra.GetNeighbors(a.X(), a.Y())
	emptyCells := make([]contracts.Cell, 0, len(neighbors))
	agents := make([]contracts.Agent, 0, len(neighbors)+1)
	agents = append(agents, a)
	for _, cell := range neighbors {
		if cell.IsEmpty() {
			emptyCells = append(emptyCells, cell)
			continue
		}
		if cell.IsAgent() {
			agents = append(agents, cell.Agent())
			continue
		}
		if cell.CellType() == enum.CellTypeWall {
			continue
		}
		panic("unknown cell type")
	}
	if len(emptyCells) == 0 {
		return nil
	}

	children, err := a.Reproduce(agents)
	if err != nil {
		return nil
	}

	placedChildren := make([]contracts.Agent, 0, len(children))

	for i := 0; i < len(children) && len(emptyCells) > 0; i++ {
		randIndex := rand.Intn(len(emptyCells))
		targetCell := emptyCells[randIndex]
		if targetCell.IsEmpty() {
			targetCell.SetAgent(children[i])
			placedChildren = append(placedChildren, children[i])
		} else {
			children[i].Kill(terra)
			children = append(children[:i], children[i+1:]...)
		}

		emptyCells = append(emptyCells[:randIndex], emptyCells[randIndex+1:]...)
	}

	return placedChildren
}

func (a *Agent) Kill(terra contracts.Terrain) {
	terra.Cell(a.X(), a.Y()).RemoveAgent()
	a.energy = 0
}
