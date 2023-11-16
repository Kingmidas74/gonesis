package nature

import (
	"errors"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Plant struct {
	contracts.ReproductionSystem

	config            *configuration.Configuration
	availableCommands []contracts.Command
}

func (a *Plant) Configure(config *configuration.Configuration, availableCommands []contracts.Command) error {
	for i := range availableCommands {
		if !availableCommands[i].IsAvailable(a) {
			return errors.New("command not available")
		}
	}
	a.config = config
	a.availableCommands = availableCommands
	return nil
}

func (a *Plant) FindCommand(commandIdentifier int) contracts.Command {
	if commandIdentifier < 0 || commandIdentifier >= len(a.availableCommands) {
		return nil
	}
	return a.availableCommands[commandIdentifier]
}

func (a *Plant) AgentType() enum.AgentType {
	return enum.AgentTypePlant
}

func (a *Plant) MaxEnergy() int {
	return a.config.PlantConfiguration.MaxEnergy
}

func (a *Plant) MaxDailyCommandCount() int {
	return a.config.PlantConfiguration.MaxDailyCommandCount
}

func (a *Plant) InitialEnergy() int {
	return a.config.PlantConfiguration.InitialEnergy
}

func (a *Plant) BrainVolume() int {
	return a.config.PlantConfiguration.BrainVolume
}

func (a *Plant) ReproductionEnergyCost() int {
	return a.config.PlantConfiguration.ReproductionEnergyCost
}

func (a *Plant) ReproductionChance() float64 {
	return a.config.PlantConfiguration.ReproductionChance
}

func (a *Plant) MutationChance() float64 {
	return a.config.PlantConfiguration.MutationChance
}

func (a *Plant) AvailableFood() map[enum.AgentType]int {
	return map[enum.AgentType]int{}
}
