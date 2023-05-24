package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"math/rand"
)

type Carnivore struct {
}

func (a Carnivore) AgentType() enum.AgentType {
	return enum.AgentTypeCarnivore
}

func (a Carnivore) Genesis(parent contracts.Agent) contracts.Agent {
	if parent.Energy() < 100 {
		return nil
	}
	if rand.Intn(100) > 80 {
		return nil
	}
	parent.IncreaseEnergy(-parent.Energy() / 2)
	return agent.NewAgent[Carnivore](parent.Energy()/2, len(parent.Commands()))
}
