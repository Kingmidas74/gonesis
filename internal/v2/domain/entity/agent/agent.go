package agent

import "github.com/google/uuid"

type Agent struct {
	brain  Brain
	nature Nature

	id     string
	energy int
}

func New(options ...func(*Agent)) *Agent {
	a := &Agent{
		id: uuid.New().String(),
	}
	for _, o := range options {
		o(a)
	}
	return a
}

func WithBrain(brain Brain) func(*Agent) {
	return func(a *Agent) {
		a.brain = brain
	}
}

func WithNature(nature Nature) func(*Agent) {
	return func(a *Agent) {
		a.nature = nature
		a.energy = nature.InitialEnergy()
	}
}

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

func (a *Agent) Action(terrain Terrain) (actionsCount int, err error) {
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
