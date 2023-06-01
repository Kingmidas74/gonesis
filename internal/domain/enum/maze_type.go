package enum

import "encoding/json"

type MazeType uint8

const (
	MazeTypeUndefined MazeType = iota
	MazeTypeAldousBroder
	MazeTypeBinary
	MazeTypeGrid
	MazeTypeBorder
	MazeTypeSideWinder
	MazeTypeEmpty
)

func (t MazeType) Value() int {
	return int(t)
}

func (t MazeType) String() string {
	switch t {
	case MazeTypeAldousBroder:
		return "aldous_broder"
	case MazeTypeBinary:
		return "binary"
	case MazeTypeGrid:
		return "grid"
	case MazeTypeBorder:
		return "border"
	case MazeTypeSideWinder:
		return "side_winder"
	case MazeTypeEmpty:
		return "empty"
	default:
		return "undefined"
	}
}

func NewMazeTypeByString(s string) MazeType {
	switch s {
	case "aldous_broder":
		return MazeTypeAldousBroder
	case "binary":
		return MazeTypeBinary
	case "grid":
		return MazeTypeGrid
	case "border":
		return MazeTypeBorder
	case "side_winder":
		return MazeTypeSideWinder
	case "empty":
		return MazeTypeEmpty
	default:
		return MazeTypeUndefined
	}
}

func (t MazeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *MazeType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*t = NewMazeTypeByString(s)

	return nil
}
