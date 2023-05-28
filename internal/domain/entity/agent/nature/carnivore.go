package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"math/rand"
)

type Carnivore struct {
}

func (a Carnivore) AgentType() enum.AgentType {
	return enum.AgentTypeCarnivore
}

func (a Carnivore) Genesis(parent contracts.Agent, config *configuration.Configuration) []contracts.Agent {
	if parent.Energy() < config.CarnivoreConfiguration.MaxEnergy {
		return nil
	}
	if rand.Intn(100) > 80 {
		return nil
	}
	parent.DecreaseEnergy(parent.Energy() / 2)
	return []contracts.Agent{agent.NewAgent[Carnivore](parent.Energy()/2, len(parent.Commands()))}
}

func (a Carnivore) MaxEnergy(config *configuration.Configuration) int {
	return config.CarnivoreConfiguration.MaxEnergy
}
