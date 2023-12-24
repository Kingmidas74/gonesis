package nature

import (
	"errors"
	contract "github.com/kingmidas74/gonesis-engine/internal/domain/contract"
)

type Nature struct {
	initialEnergy int
	initialCount  int

	brainVolume int

	maxDailyCommandCount int

	reproductionEnergyCost int
	reproductionChance     float64
	mutationChance         float64

	agentType contract.AgentType

	commands      []contract.Command
	availableFood map[contract.AgentType]int

	reproductionSystem contract.Reproduction
}

func New(reproductionSystem contract.Reproduction, options ...func(nature *Nature)) (*Nature, error) {
	b := &Nature{
		commands:           []contract.Command{},
		availableFood:      map[contract.AgentType]int{},
		reproductionSystem: reproductionSystem,
	}
	for _, o := range options {
		o(b)
	}
	if b.reproductionSystem == nil {
		return nil, ErrReproductionSystemNotDefined
	}
	return b, nil
}

func WithInitialEnergy(initialEnergy int) func(nature *Nature) {
	return func(n *Nature) {
		n.initialEnergy = initialEnergy
	}
}

func WithInitialCount(initialCount int) func(nature *Nature) {
	return func(n *Nature) {
		n.initialCount = initialCount
	}
}

func WithBrainVolume(brainVolume int) func(nature *Nature) {
	return func(n *Nature) {
		n.brainVolume = brainVolume
	}
}

func WithMaxDailyCommandCount(maxDailyCommandCount int) func(nature *Nature) {
	return func(n *Nature) {
		n.maxDailyCommandCount = maxDailyCommandCount
	}
}

func WithReproductionEnergyCost(reproductionEnergyCost int) func(nature *Nature) {
	return func(n *Nature) {
		n.reproductionEnergyCost = reproductionEnergyCost
	}
}

func WithReproductionChance(reproductionChance float64) func(nature *Nature) {
	return func(n *Nature) {
		n.reproductionChance = reproductionChance
	}
}

func WithMutationChance(mutationChance float64) func(nature *Nature) {
	return func(n *Nature) {
		n.mutationChance = mutationChance
	}
}

func WithCommands(commands []contract.Command) func(nature *Nature) {
	return func(n *Nature) {
		n.commands = commands
	}
}

func WithType(t contract.AgentType) func(nature *Nature) {
	return func(n *Nature) {
		n.agentType = t
	}
}

func WithAvailableFood(availableFood map[contract.AgentType]int) func(nature *Nature) {
	return func(n *Nature) {
		n.availableFood = availableFood
	}
}

func WithReproductionSystem(reproductionSystem contract.Reproduction) func(nature *Nature) {
	return func(n *Nature) {
		n.reproductionSystem = reproductionSystem
	}
}

var ErrReproductionSystemNotDefined = errors.New("reproduction system not defined")
