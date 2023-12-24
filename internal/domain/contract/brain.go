package contract

type Brain interface {
	// Command returns the command for the given identifier or current command if identifier is nil
	Command(identifier *int) int
	IncreaseAddress(delta int)
	Address() int
	Commands() []int
}
