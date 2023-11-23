package cell

type Coordinates interface {
	X() int
	SetX(x int)
	Y() int
	SetY(y int)
}

type Agent interface {
	IsAlive() bool
	SetX(x int)
	SetY(y int)
}

type Cell struct {
	x int
	y int

	agent Agent

	northWall bool
	southWall bool
	westWall  bool
	eastWall  bool
}

func New(x, y int) *Cell {
	return &Cell{
		x: x,
		y: y,

		northWall: true,
		southWall: true,
		westWall:  true,
		eastWall:  true,
	}
}

func (c *Cell) SetAgent(a Agent) {
	if c.agent != nil && c.agent.IsAlive() {
		panic("cell already has an agent")
	}

	c.agent = a
	a.SetX(c.x)
	a.SetY(c.y)
}

func (c *Cell) RemoveAgent() {
	c.agent = nil
}

func (c *Cell) IsEmpty() bool {
	return c.agent == nil
}

func (c *Cell) Agent() Agent {
	return c.agent
}

func (c *Cell) NorthWall() bool {
	return c.northWall
}

func (c *Cell) SetNorthWall(flag bool) {
	c.northWall = flag
}

func (c *Cell) SouthWall() bool {
	return c.southWall
}

func (c *Cell) SetSouthWall(flag bool) {
	c.southWall = flag
}

func (c *Cell) WestWall() bool {
	return c.westWall
}

func (c *Cell) SetWestWall(flag bool) {
	c.westWall = flag
}

func (c *Cell) EastWall() bool {
	return c.eastWall
}

func (c *Cell) SetEastWall(flag bool) {
	c.eastWall = flag
}
