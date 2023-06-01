package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"math/rand"
)

type Herbivore struct {
	config *configuration.Configuration
}

func (a *Herbivore) Configure(config *configuration.Configuration) {
	a.config = config
}

func (a *Herbivore) AgentType() enum.AgentType {
	return enum.AgentTypeHerbivore
}

func (a *Herbivore) Genesis(parent contracts.Agent) []contracts.Agent {
	if parent.Energy() < a.config.HerbivoreConfiguration.MaxEnergy {
		return nil
	}
	if rand.Intn(100) > 80 {
		return nil
	}
	parent.DecreaseEnergy(a.InitialEnergy())
	brain := agent.NewBrainWithCommands(parent.Commands())
	child := agent.NewAgentWithBrain(a, brain)
	return []contracts.Agent{child}
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
