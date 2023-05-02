package entity

type Coords struct {
	x int
	y int
}

func (c *Coords) X() int {
	return c.x
}

func (c *Coords) SetX(x int) {
	c.x = x
}

func (c *Coords) Y() int {
	return c.y
}

func (c *Coords) SetY(y int) {
	c.y = y
}
