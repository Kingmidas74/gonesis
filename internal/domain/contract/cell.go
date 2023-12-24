package contract

type Cell interface {
	CoordinateReader
	WallInfo

	IsEmpty() bool

	Agent() Agent
	SetAgent(agent Agent)
	RemoveAgent()
}
