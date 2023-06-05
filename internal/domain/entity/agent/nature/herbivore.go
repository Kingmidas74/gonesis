package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Herbivore struct {
	contracts.ReproductionSystem

	config *configuration.Configuration
}

func (a *Herbivore) Configure(config *configuration.Configuration) {
	a.config = config
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
