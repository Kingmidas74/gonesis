package budding

import "testing"

type agent struct {
	commands []int
}

func (a *agent) Commands() []int {
	return a.commands
}

func (a *agent) Address() int {
	return 0
}

func (a *agent) CanReproduce() bool {
	return true
}

func (a *agent) ReproductionChance() float64 {
	return 1
}

type randomFloatGenerator struct {
}

func (g randomFloatGenerator) Float64() float64 {
	return 0
}

type mutation struct {
}

func (m mutation) Mutate(subject Agent) error {
	commands := subject.Commands()
	commands[subject.Address()] = commands[subject.Address()] + 1
	return nil
}

func TestReproduction_Reproduce(t *testing.T) {
	t.Parallel()

	parent := &agent{
		commands: []int{0, 1, 2},
	}
	floatGenerator := &randomFloatGenerator{}
	mutationHandler := &mutation{}
	reproduction := New(floatGenerator, mutationHandler)

	children, err := reproduction.Reproduce([]Agent{parent})
	if err != nil {
		t.Fatal(err)
	}

	if len(children) != 1 {
		t.Fatal("invalid children count")
	}

	child := children[0]
	if len(child.Commands()) != len(parent.Commands()) {
		t.Fatal("invalid child commands count")
	}

	for i, command := range child.Commands() {
		if command != parent.Commands()[i] {
			t.Fatal("invalid child command")
		}
	}
}
