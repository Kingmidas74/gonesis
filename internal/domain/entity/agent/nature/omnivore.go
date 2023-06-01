package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"math/rand"
)

type Omnivore struct {
	config *configuration.Configuration
}

func (a *Omnivore) Configure(config *configuration.Configuration) {
	a.config = config
}

func (a *Omnivore) AgentType() enum.AgentType {
	return enum.AgentTypeOmnivore
}

func (a *Omnivore) Genesis(parent contracts.Agent) []contracts.Agent {
	if parent.Energy() < a.MaxEnergy() {
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
