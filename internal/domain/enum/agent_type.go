package enum

type AgentType uint8

const (
	AgentTypeUndefined AgentType = iota
	AgentTypeCarnivore
	AgentTypeHerbivore
	AgentTypeDecomposer
	AgentTypePlant
	AgentTypeOmnivore
	AgentTypeGround
)

func (t AgentType) Value() int {
	return int(t)
}

func (t AgentType) String() string {
	switch t {
	case AgentTypeCarnivore:
		return "carnivore"
	case AgentTypeHerbivore:
		return "herbivore"
	case AgentTypeDecomposer:
		return "decomposer"
	case AgentTypePlant:
		return "plant"
	case AgentTypeOmnivore:
		return "omnivore"
	case AgentTypeGround:
		return "ground"
	default:
		return "undefined"
	}
}

func (t AgentType) NewAgentTypeByString(s string) AgentType {
	switch s {
	case "carnivore":
		return AgentTypeCarnivore
	case "herbivore":
		return AgentTypeHerbivore
	case "decomposer":
		return AgentTypeDecomposer
	case "plant":
		return AgentTypePlant
	case "omnivore":
		return AgentTypeOmnivore
	case "ground":
		return AgentTypeGround
	default:
		return AgentTypeUndefined
	}
}
