package agent

import (
	"errors"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/brain"
)

var ErrCommandUndefined = errors.New("command is undefined")

type Agent struct {
	brain.Brain

	entity.Coords

	energy int
}

func New(energy int, brainVolume int) *Agent {
	return &Agent{
		Brain:  brain.New(brainVolume),
		energy: energy,
	}
}

func (a *Agent) Energy() int {
	return a.energy
}

func (a *Agent) SetEnergy(energy int) {
	a.energy = energy
}

func (a *Agent) IsAlive() bool {
	return a.Energy() > 0
}

// TODO: replace findCommandPredicate with []contracts.Command
func (a *Agent) NextDay(maxSteps int, terra contracts.Terrain, findCommandPredicate func(int) contracts.Command) error {
	for step := 0; a.IsAlive() && step < maxSteps; step++ {
		commandIdentifier := a.Command(nil)
		command := findCommandPredicate(commandIdentifier)
		if command == nil {
			a.SetEnergy(a.Energy() - 1)
			a.IncreaseAddress(commandIdentifier)
			return nil
		}
		delta := command.Handle(a, terra)
		a.IncreaseAddress(delta)
		if command.IsInterrupt() {
			break
		}
		a.SetEnergy(a.Energy() - 1)
	}

	return nil
}
