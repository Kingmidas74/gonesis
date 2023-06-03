package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Carnivore struct {
	contracts.ReproductionSystem

	config *configuration.Configuration
}

func (a *Carnivore) Configure(config *configuration.Configuration) {
	a.config = config
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
