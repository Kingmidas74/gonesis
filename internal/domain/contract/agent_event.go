package contract

type AgentEvent interface {
}

type AgentEventType uint8

const (
	AgentEventTypeUndefined AgentEventType = iota
)

func (t AgentEventType) Value() int {
	return int(t)
}

func (t AgentEventType) String() string {
	return "undefined"
}
