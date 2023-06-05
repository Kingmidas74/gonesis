package contracts

import "github.com/kingmidas74/gonesis-engine/internal/domain/enum"

type Mutation interface {
	MutationType() enum.MutationType
	Mutate(subject Agent) (Brain, error)
}
