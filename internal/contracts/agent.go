package contracts

type Agent interface {
	Address() int
	Command(address *int) int
	X() int
	SetX(x int)
	Y() int
	SetY(y int)
	IsAlive() bool
	NextDay(maxSteps int, world Terrain, command func(commandIdentifier int) Command) error
	Energy() int
	Commands() []int
}
