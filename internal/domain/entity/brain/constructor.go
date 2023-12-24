package brain

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

func New(options ...func(*Brain)) *Brain {
	b := &Brain{
		address:  0,
		commands: make([]int, 0),
	}
	for _, o := range options {
		o(b)
	}
	return b
}

func WithCommands(commands []int) func(brain *Brain) {
	return func(b *Brain) {
		b.commands = commands
	}
}

func WithAddress(address int) func(brain *Brain) {
	return func(b *Brain) {
		b.address = address
	}
}
