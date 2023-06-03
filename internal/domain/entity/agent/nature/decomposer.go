package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Decomposer struct {
	contracts.ReproductionSystem

	config *configuration.Configuration
}

func (a *Decomposer) Configure(config *configuration.Configuration) {
	a.config = config
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
