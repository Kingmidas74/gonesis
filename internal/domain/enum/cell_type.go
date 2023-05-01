package enum

type CellType uint8

const (
	CellTypeEmpty CellType = iota
	CellTypeLocked
	CellTypeOrganic
	CellTypeObstacle
)

func (t CellType) Value() uint8 {
	return uint8(t)
}
