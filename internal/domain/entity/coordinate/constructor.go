package coordinate

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
