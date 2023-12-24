package contract

type Nature interface {
	InitialEnergy() int
	MaxDailyCommandCount() int
	FindCommand(identifier int) Command
	Type() AgentType
	AvailableFood() map[AgentType]int
	ReproductionEnergyCost() int
	ReproductionChance() float64
	Reproduction() Reproduction
	MutationChance() float64
	InitialCount() int
}
