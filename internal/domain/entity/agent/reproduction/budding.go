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
	contracts.Mutation
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
	if parent.Energy() < parent.ReproductionEnergyCost() {
		return children, nil
	}

	if rand.Float64() < parent.ReproductionChance() {
		return children, nil
	}

	brain, err := parent.Mutate(parent)
	if err != nil {
		return children, err
	}

	parent.DecreaseEnergy(parent.ReproductionEnergyCost())
	child := agent.NewAgentWithBrain(parent, brain, parent.Generation()+1)
	return append(children, child), nil
}
