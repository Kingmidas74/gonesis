package budding

import (
	"errors"
	"github.com/kingmidas74/gonesis-engine/internal/domain/contract"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
)

func (b Reproduction) Reproduce(parents []contract.Agent) ([]contract.Agent, error) {
	children := make([]contract.Agent, 0)

	parent := parents[0]
	if parent.Energy() < parent.ReproductionEnergyCost() {
		return children, nil
	}

	if b.floatGenerator.Generate() > parent.ReproductionChance() {
		return children, nil
	}

	parentBrain := parent.Brain()

	if b.floatGenerator.Generate() < parent.MutationChance() {
		err := b.mutationHandler.Mutate(parentBrain)
		if err != nil {
			return children, err
		}
	}

	child := agent.New(
		agent.WithNature(parent.Nature()),
		agent.WithBrain(parentBrain),
		agent.WithGeneration(parent.Generation()+1),
	)
	children = append(children, child)

	return children, nil
}

var ErrInvalidParentCount = errors.New("invalid parent count")
