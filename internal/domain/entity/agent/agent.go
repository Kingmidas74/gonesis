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
	for _, cell := range terra.GetNeighbors(a.X(), a.Y()) {
		if cell.IsEmpty() {
			emptyCells = append(emptyCells, cell)
		}
	}
	if len(emptyCells) == 0 {
		return nil
	}

	targetCell := emptyCells[rand.Intn(len(emptyCells))]
	if !targetCell.IsEmpty() {
		return nil
	}

	if children := a.Genesis(a); children != nil && len(children) > 0 {
		if children[0] == nil {
			return nil
		}
		if !targetCell.IsEmpty() {
			return nil
		}
		targetCell.SetAgent(children[0])
		return children
	}

	return nil
}

func (a *Agent) Kill(terra contracts.Terrain) {
	terra.Cell(a.X(), a.Y()).RemoveAgent()
	a.energy = 0
}
