package nature

import (
	"errors"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Herbivore struct {
	contracts.ReproductionSystem

	config            *configuration.Configuration
	availableCommands []contracts.Command
}

func (a *Herbivore) Configure(config *configuration.Configuration, availableCommands []contracts.Command) error {
	for i := range availableCommands {
		if !availableCommands[i].IsAvailable(a) {
			return errors.New("command not available")
		}
	}
	a.config = config
	a.availableCommands = availableCommands
	return nil
}

func (a *Herbivore) FindCommand(commandIdentifier int) contracts.Command {
	if commandIdentifier < 0 || commandIdentifier >= len(a.availableCommands) {
		return nil
	}
	return a.availableCommands[commandIdentifier]
}

func (a *Herbivore) AgentType() enum.AgentType {
	return enum.AgentTypeHerbivore
}

func (a *Herbivore) MaxEnergy() int {
	return a.config.HerbivoreConfiguration.MaxEnergy
}

func (a *Herbivore) MaxDailyCommandCount() int {
	return a.config.HerbivoreConfiguration.MaxDailyCommandCount
}

func (a *Herbivore) InitialEnergy() int {
	return a.config.HerbivoreConfiguration.InitialEnergy
}

func (a *Herbivore) BrainVolume() int {
	return a.config.HerbivoreConfiguration.BrainVolume
}

func (a *Herbivore) ReproductionEnergyCost() int {
	return a.config.HerbivoreConfiguration.ReproductionEnergyCost
}

func (a *Herbivore) ReproductionChance() float64 {
	return a.config.HerbivoreConfiguration.ReproductionChance
}

func (a *Herbivore) MutationChance() float64 {
	return a.config.HerbivoreConfiguration.MutationChance
}

func (a *Herbivore) AvailableFood() map[enum.AgentType]int {
	return map[enum.AgentType]int{
		enum.AgentTypePlant: 16,
	}
}
