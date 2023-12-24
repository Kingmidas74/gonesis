package agent

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/contract"
)

func (a *Agent) ID() string {
	return a.id
}

func (a *Agent) IsAlive() bool {
	return a.energy > 0
}

func (a *Agent) Energy() int {
	return a.energy
}

func (a *Agent) IncreaseEnergy(delta int) {
	a.energy += delta
}

func (a *Agent) DecreaseEnergy(delta int) {
	a.energy -= delta
}

func (a *Agent) Action(terrain contract.Terrain) (actionsCount int, err error) {
	for step := 0; a.IsAlive() && step < a.nature.MaxDailyCommandCount(); step++ {
		commandIdentifier := a.brain.Command(nil)
		command := a.nature.FindCommand(commandIdentifier)
		if command == nil {
			a.brain.IncreaseAddress(commandIdentifier)
			continue
		}

		delta, err := command.Handle(a, terrain)
		if err != nil {
			return actionsCount, err
		}

		a.brain.IncreaseAddress(delta)
		actionsCount++

		if command.IsInterrupt() {
			break
		}
	}

	a.DecreaseEnergy(1)
	return actionsCount, err
}

func (a *Agent) Address() int {
	return a.brain.Address()
}

func (a *Agent) Command(address *int) int {
	return a.brain.Command(address)
}

func (a *Agent) X() int {
	return a.coordinate.X()
}

func (a *Agent) Y() int {
	return a.coordinate.Y()
}

func (a *Agent) SetX(x int) {
	a.coordinate.SetX(x)
}

func (a *Agent) SetY(y int) {
	a.coordinate.SetY(y)
}

func (a *Agent) Type() contract.AgentType {
	return a.nature.Type()
}

func (a *Agent) AvailableFood() map[contract.AgentType]int {
	return a.nature.AvailableFood()
}

func (a *Agent) Nature() contract.Nature {
	return a.nature
}

func (a *Agent) Brain() contract.Brain {
	return a.brain
}

func (a *Agent) Generation() int {
	return a.generation
}

func (a *Agent) Reproduce(terrain contract.Terrain) ([]contract.Agent, error) {
	neighbors := terrain.GetNeighbors(a.X(), a.Y())
	emptyCells := make([]contract.Cell, 0, len(neighbors))
	parents := make([]contract.Agent, 0, len(neighbors)+1)
	parents = append(parents, a)

	for _, cell := range neighbors {
		if cell.IsEmpty() {
			emptyCells = append(emptyCells, cell)
		} else if agent := cell.Agent(); agent != nil {
			parents = append(parents, agent)
		} else {
			panic("unknown cell type")
		}
	}

	if len(emptyCells) == 0 {
		return nil, nil
	}

	children, err := a.nature.Reproduction().Reproduce(parents)
	if err != nil {
		return nil, err
	}

	placedChildren := make([]contract.Agent, 0)

	for _, child := range children {
		if !child.IsAlive() {
			continue
		}
		// Find an empty cell for the child
		for i, cell := range emptyCells {
			if terrain.CanMoveTo(terrain.Cell(a.X(), a.Y()), cell) {
				cell.SetAgent(child)
				placedChildren = append(placedChildren, child)
				// Remove the used cell from emptyCells
				emptyCells = append(emptyCells[:i], emptyCells[i+1:]...)
				break
			}
		}
	}

	return placedChildren, nil
}

func (a *Agent) ReproductionEnergyCost() int {
	return a.nature.ReproductionEnergyCost()
}

func (a *Agent) ReproductionChance() float64 {
	return a.nature.ReproductionChance()
}

func (a *Agent) MutationChance() float64 {
	return a.nature.MutationChance()
}

func (a *Agent) AddEvent(event contract.AgentEvent) {
	a.events = append(a.events, event)
}

func (a *Agent) Events() []contract.AgentEvent {
	return a.events
}
