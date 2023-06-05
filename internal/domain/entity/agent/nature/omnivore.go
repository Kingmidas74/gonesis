package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Omnivore struct {
	contracts.ReproductionSystem

	config *configuration.Configuration
}

func (a *Omnivore) Configure(config *configuration.Configuration) {
	a.config = config
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
