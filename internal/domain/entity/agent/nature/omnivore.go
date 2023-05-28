package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"math/rand"
)

type Omnivore struct {
}

func (a Omnivore) AgentType() enum.AgentType {
	return enum.AgentTypeOmnivore
}

func (a Omnivore) Genesis(parent contracts.Agent, config *configuration.Configuration) []contracts.Agent {
	if parent.Energy() < config.OmnivoreConfiguration.MaxEnergy {
		return nil
	}
	if rand.Intn(100) > 80 {
		return nil
	}
	parent.DecreaseEnergy(parent.Energy() / 2)
	brain := agent.NewBrainWithCommands(parent.Commands())
	child := agent.NewAgentWithBrain[Omnivore](parent.Energy()/2, brain)
	return []contracts.Agent{child}
}

func (a Omnivore) MaxEnergy(config *configuration.Configuration) int {
	return config.OmnivoreConfiguration.MaxEnergy
}
