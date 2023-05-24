package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"math/rand"
)

type Decomposer struct {
}

func (a Decomposer) AgentType() enum.AgentType {
	return enum.AgentTypeDecomposer
}

func (a Decomposer) Genesis(parent contracts.Agent) contracts.Agent {
	if parent.Energy() < 100 {
		return nil
	}
	if rand.Intn(100) > 80 {
		return nil
	}
	parent.IncreaseEnergy(-parent.Energy() / 2)
	brain := agent.NewBrainWithCommands(parent.Commands())
	child := agent.NewAgentWithBrain[Decomposer](parent.Energy()/2, brain)
	return child
}
