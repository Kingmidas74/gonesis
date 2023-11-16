package nature

import (
	"errors"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Carnivore struct {
	contracts.ReproductionSystem

	config            *configuration.Configuration
	availableCommands []contracts.Command
}

func (a *Carnivore) Configure(config *configuration.Configuration, availableCommands []contracts.Command) error {
	for i := range availableCommands {
		if !availableCommands[i].IsAvailable(a) {
			return errors.New("command not available")
		}
	}
	a.availableCommands = availableCommands
	a.config = config
	return nil
}

func (a *Carnivore) FindCommand(commandIdentifier int) contracts.Command {
	if commandIdentifier < 0 || commandIdentifier >= len(a.availableCommands) {
		return nil
	}
	return a.availableCommands[commandIdentifier]
}

func (a *Carnivore) AgentType() enum.AgentType {
	return enum.AgentTypeCarnivore
}

func (a *Carnivore) MaxEnergy() int {
	return a.config.CarnivoreConfiguration.MaxEnergy
}

func (a *Carnivore) MaxDailyCommandCount() int {
	return a.config.CarnivoreConfiguration.MaxDailyCommandCount
}

func (a *Carnivore) InitialEnergy() int {
	return a.config.CarnivoreConfiguration.InitialEnergy
}

func (a *Carnivore) BrainVolume() int {
	return a.config.CarnivoreConfiguration.BrainVolume
}

func (a *Carnivore) ReproductionEnergyCost() int {
	return a.config.CarnivoreConfiguration.ReproductionEnergyCost
}

func (a *Carnivore) ReproductionChance() float64 {
	return a.config.CarnivoreConfiguration.ReproductionChance
}

func (a *Carnivore) MutationChance() float64 {
	return a.config.CarnivoreConfiguration.MutationChance
}

func (a *Carnivore) AvailableFood() map[enum.AgentType]int {
	return map[enum.AgentType]int{
		enum.AgentTypeHerbivore: 8,
		enum.AgentTypeOmnivore:  8,
	}
}
