package contracts

type World interface {
	Terrain

	Agents() []Agent
	Action(maxSteps int, callback func(World, int)) error
	Next() error
}
