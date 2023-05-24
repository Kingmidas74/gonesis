package enum

type CellType uint8

const (
	CellTypeUndefined CellType = iota
	CellTypeEmpty
	CellTypeWall
	CellTypeAgent
)

func (t CellType) Value() int {
	return int(t)
}

func (t CellType) String() string {
	switch t {
	case CellTypeEmpty:
		return "empty"
	case CellTypeWall:
		return "wall"
	case CellTypeAgent:
		return "agent"
	default:
		return "undefined"
	}
}

func (t CellType) NewCellTypeByString(s string) CellType {
	switch s {
	case "empty":
		return CellTypeEmpty
	case "wall":
		return CellTypeWall
	case "agent":
		return CellTypeAgent
	default:
		return CellTypeUndefined
	}
}
