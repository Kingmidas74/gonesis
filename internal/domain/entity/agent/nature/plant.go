package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"math/rand"
)

type Plant struct {
}

func (a Plant) AgentType() enum.AgentType {
	return enum.AgentTypePlant
}

func (a Plant) Genesis(parent contracts.Agent, config *configuration.AgentConfiguration) []contracts.Agent {
	if rand.Intn(100) > 90 {
		return nil
	}

	childEnergy := parent.Energy() / 10

	if parent.Energy() < config.MaxEnergy {
		parent.DecreaseEnergy(childEnergy)
		return nil
	}

	parent.DecreaseEnergy(childEnergy)
	brain := agent.NewBrainWithCommands(parent.Commands())
	child := agent.NewAgentWithBrain[Plant](childEnergy, brain)
	return []contracts.Agent{child}
}
