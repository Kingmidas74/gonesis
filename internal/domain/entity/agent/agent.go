package agent

import (
	"math/rand"

	"github.com/google/uuid"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity"
)

type Agent struct {
	contracts.Brain
	contracts.AgentNature

	entity.Coords

	id         string
	energy     int
	generation int
}

func NewAgent(nature contracts.AgentNature) contracts.Agent {
	result := &Agent{
		Brain:       NewBrain(nature.BrainVolume()),
		AgentNature: nature,
		energy:      nature.InitialEnergy(),
		generation:  0,
		id:          uuid.New().String(),
	}

	return result
}

func NewAgentWithBrain(nature contracts.AgentNature, brain contracts.Brain, generation int) contracts.Agent {
	return &Agent{
		Brain:       brain,
		AgentNature: nature,
		energy:      nature.InitialEnergy(),
		generation:  generation,
		id:          uuid.New().String(),
	}
}

func (a *Agent) ID() string {
	return a.id
}

func (a *Agent) Energy() int {
	return a.energy
}

func (a *Agent) IsAlive() bool {
	return a.Energy() > 0
}

// TODO: replace findCommandPredicate with []contracts.Command
func (a *Agent) NextDay(terra contracts.Terrain) error {
	for step := 0; a.IsAlive() && step < a.MaxDailyCommandCount(); step++ {
		commandIdentifier := a.Command(nil)
		command := a.FindCommand(commandIdentifier)
		if command == nil {
			a.IncreaseAddress(commandIdentifier)
			continue
		}

		delta := command.Handle(a, terra) //DOUBLE BECAUSE OF NOT COPY?
		a.IncreaseAddress(delta)
		if command.IsInterrupt() {
			break
		}
	}
	a.DecreaseEnergy(1)
	return nil
}

func (a *Agent) IncreaseEnergy(delta int) {
	a.energy += delta
	if a.energy > a.MaxEnergy() {
		a.energy = a.MaxEnergy()
	}
}

func (a *Agent) DecreaseEnergy(delta int) {
	a.energy -= delta
}

func (a *Agent) Generation() int {
	return a.generation
}

func (a *Agent) CreateChildren(terra contracts.Terrain, config *configuration.Configuration) []contracts.Agent {
	neighbors := terra.GetNeighbors(a.X(), a.Y())
	emptyCells := make([]contracts.Cell, 0, len(neighbors))
	agents := make([]contracts.Agent, 0, len(neighbors)+1)
	agents = append(agents, a)

	for _, cell := range neighbors {
		if cell.IsEmpty() {
			emptyCells = append(emptyCells, cell)
			continue
		}
		if cell.IsAgent() {
			agents = append(agents, cell.Agent())
			continue
		}
		panic("unknown cell type")
	}

	if len(emptyCells) == 0 {
		return nil
	}

	children, err := a.Reproduce(agents)
	if err != nil {
		return nil
	}

	placedChildren := make([]contracts.Agent, 0, len(children))
	i := 0
	for len(children) > 0 && len(emptyCells) > 0 {
		randIndex := rand.Intn(len(emptyCells))
		targetCell := emptyCells[randIndex]

		if !terra.CanMoveTo(terra.Cell(a.X(), a.Y()), targetCell) {
			i++                     // incrementing i
			if i == len(children) { // check if i is not out of bounds
				break
			}
			continue
		}

		targetCell.SetAgent(children[i])
		placedChildren = append(placedChildren, children[i])

		// Efficiently remove the item at randIndex by replacing it with the last item and shrinking the slice.
		emptyCells[randIndex] = emptyCells[len(emptyCells)-1]
		emptyCells = emptyCells[:len(emptyCells)-1]

		// Remove the child from the children slice
		children[i] = children[len(children)-1]
		children = children[:len(children)-1]
		if i == len(children) { // check if i is not out of bounds
			break
		}
	}

	return placedChildren
}

func (a *Agent) Kill(terra contracts.Terrain) {
	terra.Cell(a.X(), a.Y()).RemoveAgent()
	a.energy = 0
}
