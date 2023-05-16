package contracts

type Coords interface {
	SetX(x int)
	SetY(y int)
	X() int
	Y() int
}
