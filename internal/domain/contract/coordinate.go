package contract

type CoordinateWriter interface {
	SetX(x int)
	SetY(y int)
}

type CoordinateReader interface {
	X() int
	Y() int
}

type Coordinate interface {
	CoordinateReader
	CoordinateWriter
}
