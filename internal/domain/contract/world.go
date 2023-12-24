package contract

type World interface {
	Terrain

	Next() error
	CurrentDay() int

	PlaceAgents(agents []Agent) error
}
