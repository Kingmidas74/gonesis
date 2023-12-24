package world

import (
	contract "github.com/kingmidas74/gonesis-engine/internal/domain/contract"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"golang.org/x/exp/slices"
	"math/rand"
	"sort"
)

func (w *World) Next() error {
	handledAgents, err := w.handleAgents()
	if err != nil {
		return err
	}

	groupedAgents := make(map[contract.AgentType][]contract.Agent)

	// Group agents by type
	for _, agent := range handledAgents {
		if !agent.IsAlive() {
			continue
		}
		groupedAgents[agent.Type()] = append(groupedAgents[agent.Type()], agent)
	}

	genesisRequired := false
	for _, group := range groupedAgents {
		if len(group) <= 8 {
			genesisRequired = true
			break
		}
		maxGroupEnergy := 0
		for _, ag := range group {
			if ag.Energy() > maxGroupEnergy {
				maxGroupEnergy = ag.Energy()
			}
		}
		if maxGroupEnergy < 10 {
			genesisRequired = true
			break
		}
	}

	if !genesisRequired {
		_ = w.removeDeadAgents()

		_, err = w.populate(handledAgents)
		if err != nil {
			return err
		}
		w.currentDay++
		return nil
	}

	//plants, herbivores, carnivores, omnivores = w.extract(handledAgents)

	for _, cell := range w.Cells() {
		if cell.Agent() == nil || cell.Agent().Type() == contract.AgentTypePlant {
			continue
		}
		cell.RemoveAgent()
	}

	newAgents := make([]contract.Agent, 0)

	for _, group := range groupedAgents {
		if len(group) == 0 {
			panic("empty group")
		}
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				l := len(group) - 1
				if j < l {
					l = j
				}
				parent := group[l]
				newAgent := agent.New(
					agent.WithNature(parent.Nature()),
					agent.WithGeneration(parent.Generation()),
					agent.WithBrain(parent.Brain()),
				)
				newAgents = append(newAgents, newAgent)
			}
		}
	}

	err = w.PlaceAgents(newAgents)
	if err != nil {
		return err
	}

	w.currentDay++
	return nil
}

func (w *World) CurrentDay() int {
	return w.currentDay
}

func (w *World) Width() int {
	return w.terrain.Width()
}

func (w *World) Height() int {
	return w.terrain.Height()
}

func (w *World) Cells() []contract.Cell {
	return w.terrain.Cells()
}

func (w *World) Cell(x, y int) contract.Cell {
	return w.terrain.Cell(x, y)
}

func (w *World) EmptyCells() []contract.Cell {
	return w.terrain.EmptyCells()
}

func (w *World) GetNeighbor(x, y, direction int) contract.Cell {
	return w.terrain.GetNeighbor(x, y, direction)
}

func (w *World) CanMoveTo(from, to contract.Cell) bool {
	return w.terrain.CanMoveTo(from, to)
}

func (w *World) GetNeighbors(x, y int) []contract.Cell {
	return w.terrain.GetNeighbors(x, y)
}

func (w *World) PlaceAgents(agents []contract.Agent) error {
	emptyCells := w.terrain.EmptyCells()
	pickedCellIndexes := make([]int, 0)
	for i := 0; i < len(agents); i++ {
		targetIndex := rand.Intn(len(emptyCells))
		if slices.Contains(pickedCellIndexes, targetIndex) {
			i--
			continue
		}
		pickedCellIndexes = append(pickedCellIndexes, targetIndex)
		emptyCell := emptyCells[targetIndex]
		if agents[i] == nil {
			panic("agent is nil")
		}
		emptyCell.SetAgent(agents[i])
	}
	return nil
}

func (w *World) handleAgents() (map[string]contract.Agent, error) {
	handledAgents := make(map[string]contract.Agent)

	for _, cell := range w.terrain.Cells() {
		agent := cell.Agent()
		if agent == nil {
			continue
		}
		if _, ok := handledAgents[agent.ID()]; ok {
			continue
		}

		if _, err := agent.Action(w.terrain); err != nil {
			return nil, err
		}

		handledAgents[agent.ID()] = agent
	}

	return handledAgents, nil
}

func (w *World) extract(agents map[string]contract.Agent) ([]contract.Agent, []contract.Agent, []contract.Agent, []contract.Agent) {
	groupedAgents := make(map[contract.AgentType][]contract.Agent)
	var category1, category2, category3, category4 []contract.Agent
	// Group agents by type
	for _, agent := range agents {
		if !agent.IsAlive() {
			continue
		}
		groupedAgents[agent.Type()] = append(groupedAgents[agent.Type()], agent)
	}

	// Sort each group and take the top 8 agents
	for _, group := range groupedAgents {
		// Sorting by generation, then by energy
		sort.Slice(group, func(i, j int) bool {
			if group[i].Generation() == group[j].Generation() {
				return group[i].Energy() > group[j].Energy()
			}
			return group[i].Generation() < group[j].Generation()
		})

		// Take the top 8 agents
		if len(group) > 8 {
			group = group[:8]
		}

		// Assign to respective category
		switch group[0].Type().Value() {
		case 1:
			category1 = group
		case 2:
			category2 = group
		case 3:
			category3 = group
		case 4:
			category4 = group
		}
	}

	return category1, category2, category3, category4
}

func (w *World) populate(handledAgents map[string]contract.Agent) ([]contract.Agent, error) {
	agents := make([]contract.Agent, 0)
	for _, agent := range handledAgents {
		if !agent.IsAlive() {
			continue
		}

		children, err := agent.Reproduce(w.terrain)
		if err != nil {
			return nil, err
		}
		agents = append(agents, children...)
	}

	return agents, nil
}

func (w *World) Agents() []contract.Agent {
	agents := make([]contract.Agent, 0)
	for _, cell := range w.terrain.Cells() {
		if agent := cell.Agent(); agent != nil {
			agents = append(agents, agent)
		}
	}
	return agents
}

func (w *World) removeDeadAgents() []contract.Agent {
	deadAgents := make([]contract.Agent, 0)
	for _, cell := range w.Cells() {
		if agent := cell.Agent(); agent != nil && !agent.IsAlive() {
			deadAgents = append(deadAgents, agent)
			cell.RemoveAgent()
		}
	}
	return deadAgents
}
