package enum

import "encoding/json"

type ReproductionSystemType uint8

const (
	ReproductionSystemTypeUndefined ReproductionSystemType = iota
	ReproductionSystemTypeBudding
)

func (t ReproductionSystemType) Value() int {
	return int(t)
}

func (t ReproductionSystemType) String() string {
	switch t {
	case ReproductionSystemTypeBudding:
		return "budding"
	default:
		return "undefined"
	}
}

func (t ReproductionSystemType) NewReproductionSystemTypeByString(s string) ReproductionSystemType {
	switch s {
	case "budding":
		return ReproductionSystemTypeBudding
	default:
		return ReproductionSystemTypeUndefined
	}
}

func (t ReproductionSystemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *ReproductionSystemType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*t = t.NewReproductionSystemTypeByString(s)

	return nil
}
