package enum

import "encoding/json"

type MutationType uint8

const (
	MutationTypeUndefined MutationType = iota
	MutationTypeRandomize
)

func (t MutationType) Value() int {
	return int(t)
}

func (t MutationType) String() string {
	switch t {
	case MutationTypeRandomize:
		return "randomize"
	default:
		return "undefined"
	}
}

func (t MutationType) NewMutationTypeByString(s string) MutationType {
	switch s {
	case "randomize":
		return MutationTypeRandomize
	default:
		return MutationTypeUndefined
	}
}

func (t MutationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *MutationType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*t = t.NewMutationTypeByString(s)

	return nil
}
