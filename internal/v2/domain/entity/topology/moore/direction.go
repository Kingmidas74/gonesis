package moore

type Direction uint8

const (
	DirectionUp Direction = iota
	DirectionUpRight
	DirectionRight
	DirectionRightDown
	DirectionDown
	DirectionDownLeft
	DirectionLeft
	DirectionLeftUp
)

func (d Direction) Value() uint8 {
	return uint8(d)
}
