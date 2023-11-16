package contracts

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type AgentNature interface {
	ReproductionSystem

	Configure(config *configuration.Configuration, availableCommands []Command) error
	FindCommand(commandIdentifier int) Command

	AgentType() enum.AgentType
	AvailableFood() map[enum.AgentType]int
	MaxEnergy() int
	MaxDailyCommandCount() int
	InitialEnergy() int
	BrainVolume() int

	ReproductionEnergyCost() int
	ReproductionChance() float64
	MutationChance() float64
}
