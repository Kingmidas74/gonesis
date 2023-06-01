package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"math/rand"
)

type Carnivore struct {
	config *configuration.Configuration
}

func (a *Carnivore) Configure(config *configuration.Configuration) {
	a.config = config
}

func (a *Carnivore) AgentType() enum.AgentType {
	return enum.AgentTypeCarnivore
}

func (a *Carnivore) Genesis(parent contracts.Agent) []contracts.Agent {
	if parent.Energy() < a.MaxEnergy() {
		return nil
	}
	if rand.Intn(100) > 80 {
		return nil
	}
	parent.DecreaseEnergy(a.InitialEnergy())
	return []contracts.Agent{agent.NewAgent(a)}
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
