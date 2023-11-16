package contracts

type Command interface {
	Handle(agent Agent, terra Terrain) (delta int)
	IsInterrupt() bool
	IsAvailable(agent AgentNature) bool
}
