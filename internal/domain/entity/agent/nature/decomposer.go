package nature

import (
	"errors"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Decomposer struct {
	contracts.ReproductionSystem

	config            *configuration.Configuration
	availableCommands []contracts.Command
}

func (a *Decomposer) Configure(config *configuration.Configuration, availableCommands []contracts.Command) error {
	for i := range availableCommands {
		if !availableCommands[i].IsAvailable(a) {
			return errors.New("command not available")
		}
	}
	a.config = config
	a.availableCommands = availableCommands
	return nil
}

func (a *Decomposer) FindCommand(commandIdentifier int) contracts.Command {
	if commandIdentifier < 0 || commandIdentifier >= len(a.availableCommands) {
		return nil
	}
	return a.availableCommands[commandIdentifier]
}

func (a *Decomposer) AgentType() enum.AgentType {
	return enum.AgentTypeDecomposer
}

func (a *Decomposer) MaxEnergy() int {
	return a.config.DecomposerConfiguration.MaxEnergy
}

func (a *Decomposer) MaxDailyCommandCount() int {
	return a.config.DecomposerConfiguration.MaxDailyCommandCount
}

func (a *Decomposer) InitialEnergy() int {
	return a.config.DecomposerConfiguration.InitialEnergy
}

func (a *Decomposer) BrainVolume() int {
	return a.config.DecomposerConfiguration.BrainVolume
}

func (a *Decomposer) ReproductionEnergyCost() int {
	return a.config.DecomposerConfiguration.ReproductionEnergyCost
}

func (a *Decomposer) ReproductionChance() float64 {
	return a.config.DecomposerConfiguration.ReproductionChance
}

func (a *Decomposer) MutationChance() float64 {
	return a.config.DecomposerConfiguration.MutationChance
}

func (a *Decomposer) AvailableFood() map[enum.AgentType]int {
	return map[enum.AgentType]int{}
}
