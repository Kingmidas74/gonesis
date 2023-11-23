package agent

type Command interface {
	Handle(agent *Agent, terrain Terrain) (int, error)
	IsInterrupt() bool
}

type Brain interface {
	Command(identifier *int) int
	IncreaseAddress(delta int)
}

type Nature interface {
	MaxDailyCommandCount() int
	FindCommand(identifier int) Command
	InitialEnergy() int
}

type Terrain interface {
}
