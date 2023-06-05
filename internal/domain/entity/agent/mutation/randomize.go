package mutation

import (
	"math/rand"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type RandomizeMutation struct {
}

func (b RandomizeMutation) MutationType() enum.MutationType {
	return enum.MutationTypeRandomize
}

func (b RandomizeMutation) Mutate(subject contracts.Agent) (contracts.Brain, error) {
	if rand.Float64() < subject.MutationChance() {
		return subject, nil
	}

	commands := subject.Commands()
	if len(commands) == 0 {
		return subject, nil
	}

	randCommandIndex := rand.Intn(len(commands))
	randNewCommand := rand.Intn(len(commands))
	commands[randCommandIndex] = randNewCommand
	return agent.NewBrainWithCommands(commands), nil
}
