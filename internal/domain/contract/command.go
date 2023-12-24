package contract

type Command interface {
	IsInterrupt() bool
	Handle(agent Agent, terra Terrain) (int, error)
}
