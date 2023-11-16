package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Ground struct {
	contracts.ReproductionSystem

	config            *configuration.Configuration
	availableCommands []contracts.Command
}

func (a *Ground) Configure(config *configuration.Configuration, availableCommands []contracts.Command) error {
	a.config = config
	a.availableCommands = availableCommands
	return nil
}

func (a *Ground) FindCommand(commandIdentifier int) contracts.Command {
	return nil
}

func (a *Ground) AgentType() enum.AgentType {
	return enum.AgentTypeGround
}

func (a *Ground) MaxEnergy() int {
	return 0
}

func (a *Ground) BrainVolume() int {
	return 0
}

func (a *Ground) MaxDailyCommandCount() int {
	return 0
}

func (a *Ground) InitialEnergy() int {
	return 0
}

func (a *Ground) ReproductionEnergyCost() int {
	return 0
}

func (a *Ground) ReproductionChance() float64 {
	return 0
}

func (a *Ground) MutationChance() float64 {
	return 0
}

func (a *Ground) AvailableFood() map[enum.AgentType]int {
	return map[enum.AgentType]int{}
}
