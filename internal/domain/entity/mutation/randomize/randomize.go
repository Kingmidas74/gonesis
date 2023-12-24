package randomize

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/contract"
)

func (b Mutation) Mutate(subject contract.Brain) error {
	commands := subject.Commands()
	if len(commands) <= 0 {
		return nil
	}

	commands[subject.Address()] = b.randomIntGenerator.Generate(len(commands))
	return nil
}
