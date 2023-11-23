package eat

type Agent interface {
	Address() int
	Command(address *int) int
	X() int
	Y() int
	IsAlive() bool
	Energy() int
	IncreaseEnergy(energy int)
	DecreaseEnergy(energy int)
}

type CellType interface {
	Value() int
}

type Cell interface {
	CellType() CellType
	SetAgent(agent Agent)
	RemoveAgent()
	Agent() Agent
}

type Terrain interface {
	GetNeighbor(x, y, direction int) Cell
	Cell(x, y int) Cell
	CanMoveTo(from, to Cell) bool
}

type Command struct {
	isInterrupt bool
}

func New(isInterrupt bool) *Command {
	return &Command{
		isInterrupt: isInterrupt,
	}
}

func (c *Command) IsInterrupt() bool {
	return c.isInterrupt
}

func (c *Command) Handle(agent Agent, terra Terrain) int {
	whereAddress := agent.Address() + 1
	direction := agent.Command(&whereAddress)
	targetCell := terra.GetNeighbor(agent.X(), agent.Y(), direction)
	if targetCell == nil {
		return 1
	}

	targetAgent := targetCell.Agent()
	if targetAgent == nil {
		return 1
	}

	if targetCell.Agent().IsAlive() {
		targetEnergy := targetCell.Agent().Energy()

		agent.IncreaseEnergy(targetEnergy)
		targetCell.Agent().DecreaseEnergy(targetEnergy)
		targetCell.RemoveAgent()
		targetCell.SetAgent(agent)
	}

	localDelta := targetCell.CellType().Value() + 1
	deltaAddress := agent.Address() + localDelta
	delta := agent.Command(&deltaAddress)
	return delta
}
