package contracts

type World interface {
	Terrain

	Agents() []Agent
	Next() error
}
