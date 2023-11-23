package neumann

type Direction uint8

const (
	DirectionUp Direction = iota
	DirectionRight
	DirectionDown
	DirectionLeft
)

func (d Direction) Value() uint8 {
	return uint8(d)
}
