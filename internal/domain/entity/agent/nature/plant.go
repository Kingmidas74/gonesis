package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Plant struct {
	contracts.ReproductionSystem

	config *configuration.Configuration
}

func (a *Plant) Configure(config *configuration.Configuration) {
	a.config = config
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
