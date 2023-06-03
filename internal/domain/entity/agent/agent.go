package agent

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"math/rand"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity"
)

type Agent struct {
	Brain

	entity.Coords

	contracts.AgentNature

	energy int

	config *configuration.Configuration
}

func NewAgent(nature contracts.AgentNature) contracts.Agent {
	return &Agent{
		Brain:       NewBrain(nature.BrainVolume()),
		AgentNature: nature,
		energy:      nature.InitialEnergy(),
	}
}

func NewAgentWithBrain(nature contracts.AgentNature, brain Brain) contracts.Agent {
	return &Agent{
		Brain:       brain,
		AgentNature: nature,
		energy:      nature.InitialEnergy(),
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
	if a.energy > 500 {
		a.energy = 500
	}
}

func (a *Agent) DecreaseEnergy(delta int) {
	a.energy -= delta
}

func (a *Agent) CreateChildren(terra contracts.Terrain, config *configuration.Configuration) []contracts.Agent {
	emptyCells := make([]contracts.Cell, 0)
	agents := make([]contracts.Agent, 0)
	agents = append(agents, a)
	for _, cell := range terra.GetNeighbors(a.X(), a.Y()) {
		if cell.IsEmpty() {
			emptyCells = append(emptyCells, cell)
		}
		if cell.IsAgent() {
			agents = append(agents, cell.Agent())
		}
	}
	if len(emptyCells) == 0 {
		return nil
	}

	children, err := a.Reproduce(agents)
	if err != nil {
		return nil
	}

	placedChildren := make([]contracts.Agent, 0)

	for i := 0; i < len(children) && len(emptyCells) > 0; i++ {
		randIndex := rand.Intn(len(emptyCells))
		targetCell := emptyCells[randIndex]
		if targetCell.IsEmpty() {
			targetCell.SetAgent(children[i])
			placedChildren = append(placedChildren, children[i])
		}
		emptyCells = append(emptyCells[:randIndex], emptyCells[randIndex+1:]...)
	}

	return placedChildren
}

func (a *Agent) Kill(terra contracts.Terrain) {
	terra.Cell(a.X(), a.Y()).RemoveAgent()
	a.energy = 0
}
