package brain

import (
	"math/rand"

	"github.com/kingmidas74/gonesis-engine/internal/util"
)

type Brain struct {
	commands []int
	address  int
}

func New(availableCommands []int, volume int) Brain {
	return Brain{
		address:  0,
		commands: generateCommandsSequence(availableCommands, volume),
	}
}

func (b *Brain) IncreaseAddress(delta int) {
	b.SetAddress(b.address + delta)
}

func (b *Brain) SetAddress(address int) {
	b.address = b.mod(address)
}

func (b *Brain) Address() int {
	return b.address
}

func (b *Brain) Command(identifier *int) int {
	address := b.address
	if identifier != nil {
		address = *identifier
	}
	return b.commands[b.mod(address)]
}

func (b *Brain) Commands() []int {
	return b.commands
}

func (b *Brain) mod(address int) int {
	return util.ModLikePython(address, len(b.commands))
}

func generateCommandsSequence(availableCommands []int, sequenceLength int) []int {
	return []int{0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 3, 0, 2, 0, 2, 2, 0, 2}
	result := make([]int, sequenceLength)
	for i := 0; i < sequenceLength; i++ {
		index := rand.Intn(len(availableCommands))
		result[i] = availableCommands[index]
	}
	return result
}
