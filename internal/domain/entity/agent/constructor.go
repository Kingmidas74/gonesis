package agent

import (
	"github.com/google/uuid"
	contract "github.com/kingmidas74/gonesis-engine/internal/domain/contract"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/coordinate"
)

type Agent struct {
	brain      contract.Brain
	nature     contract.Nature
	coordinate contract.Coordinate

	id         string
	energy     int
	generation int

	events []contract.AgentEvent
}

func New(options ...func(*Agent)) *Agent {
	a := &Agent{
		id:         uuid.New().String(),
		coordinate: coordinate.New(0, 0),
		generation: 0,
	}
	for _, o := range options {
		o(a)
	}
	return a
}

func WithBrain(brain contract.Brain) func(*Agent) {
	return func(a *Agent) {
		a.brain = brain
	}
}

func WithNature(nature contract.Nature) func(*Agent) {
	return func(a *Agent) {
		a.nature = nature
		a.energy = nature.InitialEnergy()
	}
}

func WithGeneration(generation int) func(*Agent) {
	return func(a *Agent) {
		a.generation = generation
	}
}
