package nature

type Agent interface {
	IsAlive() bool
	Energy() int
	IncreaseEnergy(delta int)
	DecreaseEnergy(delta int)
	Action(terrain Terrain) (actionsCount int, err error)
}

type Cell interface {
	Agent() Agent
	RemoveAgent()
}

type Terrain interface {
	Cells() []Cell
}

type Command interface {
	Handle(agent Agent, terrain Terrain) (int, error)
	IsInterrupt() bool
}

type Nature struct {
	initialEnergy int

	brainVolume int

	maxDailyCommandCount int

	reproductionEnergyCost int
	reproductionChance     int
	mutationChance         int

	commands []Command
}

func New() *Nature {
	return &Nature{}
}

func (n *Nature) BrainVolume() int {
	return n.brainVolume
}

func (n *Nature) ReproductionEnergyCost() int {
	return n.reproductionEnergyCost
}

func (n *Nature) ReproductionChance() int {
	return n.reproductionChance
}

func (n *Nature) MutationChance() int {
	return n.mutationChance
}

func (n *Nature) MaxDailyCommandCount() int {
	return n.maxDailyCommandCount
}

func (n *Nature) FindCommand(identifier int) Command {
	if identifier < 0 || identifier >= len(n.commands) {
		return nil
	}
	return n.commands[identifier]
}

func (n *Nature) InitialEnergy() int {
	return n.initialEnergy
}
