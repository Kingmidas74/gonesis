package budding

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/contract"
)

type Mutation interface {
	Mutate(subject contract.Brain) error
}

type RandomFloatGenerator interface {
	Generate() float64
}

type Reproduction struct {
	floatGenerator  RandomFloatGenerator
	mutationHandler Mutation
}

func New(floatGenerator RandomFloatGenerator, mutationHandler Mutation) *Reproduction {
	return &Reproduction{
		floatGenerator:  floatGenerator,
		mutationHandler: mutationHandler,
	}
}
