package budding

import (
	"errors"
)

type Agent interface {
	Commands() []int
	Address() int

	CanReproduce() bool
	ReproductionChance() float64
}

type Mutation interface {
	Mutate(subject Agent) error
}

type RandomFloatGenerator interface {
	Float64() float64
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

func (b Reproduction) Reproduce(parents []Agent) ([]Agent, error) {
	children := make([]Agent, 0)

	if len(parents) != 1 {
		return children, ErrInvalidParentCount
	}

	parent := parents[0]
	if !parent.CanReproduce() {
		return children, nil
	}

	if b.floatGenerator.Float64() > parent.ReproductionChance() {
		return children, nil
	}

	child := parent

	err := b.mutationHandler.Mutate(child)
	if err != nil {
		return children, err
	}
	children = append(children, child)

	return children, nil
}

var ErrInvalidParentCount = errors.New("invalid parent count")
