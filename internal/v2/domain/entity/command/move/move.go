package move

type Agent interface {
	Address() int
	Command(address *int) int
	X() int
	Y() int
}

type CellType interface {
	Value() int
}

type Cell interface {
	CellType() CellType
	SetAgent(agent Agent)
	RemoveAgent()
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
		panic("target cell is nil")
	}

	originalTargetCellType := targetCell.CellType()
	if terra.CanMoveTo(terra.Cell(agent.X(), agent.Y()), targetCell) {
		terra.Cell(agent.X(), agent.Y()).RemoveAgent()
		targetCell.SetAgent(agent)
	}

	localDelta := originalTargetCellType.Value() + 1
	deltaAddress := agent.Address() + localDelta
	delta := agent.Command(&deltaAddress)
	return delta
}
