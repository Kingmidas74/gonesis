package nature

import (
	"errors"
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Omnivore struct {
	contracts.ReproductionSystem

	config            *configuration.Configuration
	availableCommands []contracts.Command
}

func (a *Omnivore) Configure(config *configuration.Configuration, availableCommands []contracts.Command) error {
	for i := range availableCommands {
		if !availableCommands[i].IsAvailable(a) {
			return errors.New("command not available")
		}
	}
	a.availableCommands = availableCommands
	a.config = config
	return nil
}

func (a *Omnivore) FindCommand(commandIdentifier int) contracts.Command {
	if commandIdentifier < 0 || commandIdentifier >= len(a.availableCommands) {
		return nil
	}
	return a.availableCommands[commandIdentifier]
}

func (a *Omnivore) AgentType() enum.AgentType {
	return enum.AgentTypeOmnivore
}

func (a *Omnivore) MaxEnergy() int {
	return a.config.OmnivoreConfiguration.MaxEnergy
}

func (a *Omnivore) MaxDailyCommandCount() int {
	return a.config.OmnivoreConfiguration.MaxDailyCommandCount
}

func (a *Omnivore) InitialEnergy() int {
	return a.config.OmnivoreConfiguration.InitialEnergy
}

func (a *Omnivore) BrainVolume() int {
	return a.config.OmnivoreConfiguration.BrainVolume
}

func (a *Omnivore) ReproductionEnergyCost() int {
	return a.config.OmnivoreConfiguration.ReproductionEnergyCost
}

func (a *Omnivore) ReproductionChance() float64 {
	return a.config.OmnivoreConfiguration.ReproductionChance
}

func (a *Omnivore) MutationChance() float64 {
	return a.config.OmnivoreConfiguration.MutationChance
}

func (a *Omnivore) AvailableFood() map[enum.AgentType]int {
	return map[enum.AgentType]int{
		enum.AgentTypePlant:     1,
		enum.AgentTypeHerbivore: 4,
		enum.AgentTypeCarnivore: 4,
		enum.AgentTypeOmnivore:  2,
	}
}
