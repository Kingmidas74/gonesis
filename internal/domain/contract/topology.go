package contract

import "encoding/json"

type Topology interface {
	GetNeighbor(x, y int, direction int) Coordinate
	GetNeighbors(x, y int) []Coordinate
	CanMoveTo(currentCell, targetCell Cell, terra Terrain) bool
}

type TopologyType uint8

const (
	TopologyTypeUndefined TopologyType = iota
	TopologyTypeNeumann
	TopologyTypeMoore
)

func (t TopologyType) Value() int {
	return int(t)
}

func (t TopologyType) String() string {
	switch t {
	case TopologyTypeNeumann:
		return "neumann"
	case TopologyTypeMoore:
		return "moore"
	default:
		return "undefined"
	}
}

func (t TopologyType) NewTopologyTypeByString(s string) TopologyType {
	switch s {
	case "neumann":
		return TopologyTypeNeumann
	case "moore":
		return TopologyTypeMoore
	default:
		return TopologyTypeUndefined
	}
}

func (t TopologyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *TopologyType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*t = t.NewTopologyTypeByString(s)

	return nil
}
