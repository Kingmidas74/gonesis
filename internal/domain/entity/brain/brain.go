package brain

import "github.com/kingmidas74/gonesis-engine/internal/util"

type Brain struct {
	commands       []int
	currentAddress int
}

func New(commands []int) Brain {
	return Brain{
		currentAddress: 0,
		commands:       commands,
	}
}

func (b *Brain) mod(address int) int {
	return util.ModLikePython(address, len(b.commands))
}

func (b *Brain) MoveAddressOn(delta int) {
	b.SetCurrentAddress(b.currentAddress + delta)
}

func (b *Brain) SetCurrentAddress(address int) {
	b.currentAddress = b.mod(address)
}

func (b *Brain) CurrentAddress() int {
	return b.currentAddress
}

func (b *Brain) CurrentCommandIdentifier() int {
	return b.commands[b.mod(b.currentAddress)]
}
