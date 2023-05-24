package agent

import (
	"github.com/kingmidas74/gonesis-engine/internal/util"
	"math/rand"
)

type Brain struct {
	commands []int
	address  int
}

func NewBrain(volume int) Brain {
	return Brain{
		address:  0,
		commands: generateCommandsSequence(volume),
	}
}

func NewBrainWithCommands(commands []int) Brain {
	return Brain{
		commands: commands,
		address:  0,
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

func generatePSOnly(sequenceLength int) []int {
	result := make([]int, sequenceLength)
	for i := 0; i < sequenceLength; i++ {
		result[i] = 2
	}
	return result
}

func generateCommandsSequence(sequenceLength int) []int {
	// Create weighted buckets
	buckets := make([]int, 0)
	for i := 0; i < sequenceLength; i++ {
		weight := sequenceLength - i // calculate the weight for each number
		for j := 0; j < weight; j++ {
			buckets = append(buckets, i)
		}
	}

	// Shuffle the bucket to randomize the order
	rand.Shuffle(len(buckets), func(i, j int) {
		buckets[i], buckets[j] = buckets[j], buckets[i]
	})

	// Trim the buckets to the required sequence length
	buckets = buckets[:sequenceLength]

	return buckets
}
