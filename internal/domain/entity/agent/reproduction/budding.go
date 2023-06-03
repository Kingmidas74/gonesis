package reproduction

import (
	"errors"
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"math/rand"
)

var ErrInvalidParentCount = errors.New("invalid parent count")

type BuddingReproduction struct {
}

func (b BuddingReproduction) ReproductionType() enum.ReproductionSystemType {
	return enum.ReproductionSystemTypeBudding
}

func (b BuddingReproduction) Reproduce(parents []contracts.Agent) ([]contracts.Agent, error) {
	children := make([]contracts.Agent, 0)
	if len(parents) == 0 {
		return children, ErrInvalidParentCount
	}

	parent := parents[0]
	if parent.Energy() < parent.MaxEnergy() {
		return children, nil
	}

	if rand.Intn(100) > 90 {
		return children, nil
	}

	parent.DecreaseEnergy(parent.InitialEnergy())
	brain := agent.NewBrainWithCommands(parent.Commands())
	child := agent.NewAgentWithBrain(parent, brain)
	return append(children, child), nil
}
