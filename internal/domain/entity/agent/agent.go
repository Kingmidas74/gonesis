package agent

import (
	"errors"

	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/brain"
)

var ErrCommandUndefined = errors.New("command is undefined")

type Command interface {
	Handle(agent *Agent) (delta int)
	IsInterrupt() bool
}

type Agent struct {
	brain.Brain

	energy int
}

func New(energy int, commands []int) *Agent {
	return &Agent{
		Brain:  brain.New(commands),
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
	return a.energy > 0
}

func (a *Agent) NextDay(maxSteps int, findCommandPredicate func(int) Command) error {
	for step := 0; a.IsAlive() && step < maxSteps; step++ {
		commandIdentifier := a.CurrentCommandIdentifier()
		command := findCommandPredicate(commandIdentifier)
		if command == nil {
			return ErrCommandUndefined
		}
		delta := command.Handle(a)
		a.MoveAddressOn(delta)
		if command.IsInterrupt() {
			break
		}
		a.energy--
	}

	return nil
}
