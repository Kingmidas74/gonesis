package brain

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/util"
	"math"
)

type Brain struct {
	commands        []int
	address         int
	subroutineStack []subroutine
}

type subroutine struct {
	start         int
	count         int
	returnAddress int
}

func New() *Brain {
	return &Brain{
		address:  0,
		commands: make([]int, 0),
	}
}

func (b *Brain) IncreaseAddress(delta int) {
	b.setAddress(b.address + delta)
}

func (b *Brain) Commands() []int {
	return b.commands
}

func (b *Brain) Command(identifier *int) int {
	address := b.address
	if identifier != nil {
		address = *identifier
	}
	return b.commands[b.mod(address, len(b.commands))]
}

func (b *Brain) Address() int {
	return b.address
}

func (b *Brain) setAddress(address int) {
	if len(b.subroutineStack) > 0 {
		sub := b.subroutineStack[len(b.subroutineStack)-1]
		b.address = b.mod(b.mod(address, sub.count)+sub.start, len(b.commands))
		return
	}

	if len(b.commands) > 0 {
		b.address = b.mod(address, len(b.commands))
	}
}

func (b *Brain) mod(address, length int) int {
	return util.ModLikePython(address, length)
}

func (b *Brain) KeepAddress(from, count int) {
	if len(b.subroutineStack) > 1 {
		return
	}
	quarter := len(b.commands) / 4
	sub := subroutine{
		start:         from,
		count:         quarter,
		returnAddress: b.address,
	}
	b.subroutineStack = append(b.subroutineStack, sub)
	b.address = sub.start
}

func (b *Brain) Return() {
	if len(b.subroutineStack) == 0 {
		return
	}

	sub := b.subroutineStack[len(b.subroutineStack)-1]
	b.setAddress(sub.returnAddress + 1)
	b.subroutineStack = b.subroutineStack[:len(b.subroutineStack)-1]
}

func (b *Brain) Equals(other contracts.Brain) (difference int) {
	difference = 0

	if len(b.commands) != len(other.Commands()) {
		difference = int(math.Abs(float64(len(b.commands) - len(other.Commands()))))
	}

	for i := 0; i < len(b.commands); i++ {
		if b.commands[i] != other.Commands()[i] {
			difference++
		}
	}
	return difference
}
