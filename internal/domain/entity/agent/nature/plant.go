package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"math/rand"
)

type Plant struct {
	config *configuration.Configuration
}

func (a *Plant) Configure(config *configuration.Configuration) {
	a.config = config
}

func (a *Plant) AgentType() enum.AgentType {
	return enum.AgentTypePlant
}

func (a *Plant) Genesis(parent contracts.Agent) []contracts.Agent {
	if rand.Intn(100) > 90 {
		return nil
	}

	childEnergy := a.InitialEnergy()

	if parent.Energy() < a.MaxEnergy() {
		parent.DecreaseEnergy(2)
		return nil
	}

	parent.DecreaseEnergy(childEnergy)
	brain := agent.NewBrainWithCommands(parent.Commands())
	child := agent.NewAgentWithBrain(a, brain)
	return []contracts.Agent{child}
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
