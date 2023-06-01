package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"math/rand"
)

type Decomposer struct {
	config *configuration.Configuration
}

func (a *Decomposer) Configure(config *configuration.Configuration) {
	a.config = config
}

func (a *Decomposer) AgentType() enum.AgentType {
	return enum.AgentTypeDecomposer
}

func (a *Decomposer) Genesis(parent contracts.Agent) []contracts.Agent {
	if parent.Energy() < a.MaxEnergy() {
		return nil
	}
	if rand.Intn(100) > 80 {
		return nil
	}
	parent.DecreaseEnergy(a.InitialEnergy())
	return []contracts.Agent{agent.NewAgent(a)}
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
