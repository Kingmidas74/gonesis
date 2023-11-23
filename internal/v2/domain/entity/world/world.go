package world

type Agent interface {
	ID() string
	Action(terrain Terrain) (actionsCount int, err error)
	IsAlive() bool
}

type Cell interface {
	Agent() Agent
	RemoveAgent()
}

type Terrain interface {
	Cells() []Cell
}

type World struct {
	terrain Terrain

	currentDay int
}

func New(options ...func(*World)) *World {
	w := &World{
		currentDay: 0,
	}
	for _, o := range options {
		o(w)
	}
	return w
}

func WithTerrain(terrain Terrain) func(*World) {
	return func(w *World) {
		w.terrain = terrain
	}
}

func (w *World) NextDay() error {
	handledAgents := make(map[string]Agent)
	for _, cell := range w.terrain.Cells() {
		agent := cell.Agent()
		if agent == nil {
			continue
		}

		if _, ok := handledAgents[agent.ID()]; ok {
			continue
		}

		if _, err := agent.Action(w.terrain); err != nil {
			return err
		}

		if !agent.IsAlive() {
			cell.RemoveAgent()
			continue
		}

		handledAgents[agent.ID()] = agent
	}

	w.currentDay++
	return nil
}

func (w *World) CurrentDay() int {
	return w.currentDay
}
