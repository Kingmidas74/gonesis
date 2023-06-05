package contracts

type Brain interface {
	IncreaseAddress(delta int)
	SetAddress(address int)
	Address() int
	Command(identifier *int) int
	Commands() []int
}
