package brain

import "github.com/kingmidas74/gonesis-engine/internal/util"

type Brain struct {
	commands []int
	address  int
}

func New(commands []int) Brain {
	return Brain{
		address:  0,
		commands: commands,
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

func (b *Brain) mod(address int) int {
	return util.ModLikePython(address, len(b.commands))
}
