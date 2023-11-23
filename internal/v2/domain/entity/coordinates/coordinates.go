package coordinates

type Coordinates struct {
	x int
	y int
}

func New(x, y int) *Coordinates {
	return &Coordinates{
		x: x,
		y: y,
	}
}

func (c *Coordinates) X() int {
	return c.x
}

func (c *Coordinates) SetX(x int) {
	c.x = x
}

func (c *Coordinates) Y() int {
	return c.y
}

func (c *Coordinates) SetY(y int) {
	c.y = y
}
