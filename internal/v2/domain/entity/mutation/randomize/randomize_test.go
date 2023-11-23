package randomize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type randomIntGenerator struct {
}

func (r randomIntGenerator) Intn(n int) int {
	if n <= 0 {
		return 0
	}

	return n - 1
}

type brain struct {
	commands []int
}

func (b *brain) Commands() []int {
	return b.commands
}

func (b *brain) Address() int {
	return 1
}

func TestMutation_Mutate(t *testing.T) {
	t.Parallel()

	subject := &brain{
		commands: []int{0, 1, 2},
	}
	generator := &randomIntGenerator{}
	mutation := New(generator)

	err := mutation.Mutate(subject)
	assert.NoError(t, err)
	assert.Equal(t, []int{0, 2, 2}, subject.Commands())
}
