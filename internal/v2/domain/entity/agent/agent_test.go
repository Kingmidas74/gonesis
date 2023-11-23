package agent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type brainMock struct {
	commands []int
	address  int
}

func (b *brainMock) Command(identifier *int) int {
	if identifier == nil {
		return b.commands[b.address]
	}
	return b.commands[*identifier]
}

func (b *brainMock) IncreaseAddress(delta int) {
}

type commandMock struct {
	isInterrupt bool
	delta       int
}

func (c *commandMock) Handle(agent *Agent, terrain Terrain) (delta int, err error) {
	return c.delta, nil
}

func (c *commandMock) IsInterrupt() bool {
	return c.isInterrupt
}

type natureMock struct {
	maxDailyCommandCount int
	initialEnergy        int
	command              Command
}

func (n *natureMock) MaxDailyCommandCount() int {
	return n.maxDailyCommandCount
}

func (n *natureMock) FindCommand(identifier int) Command {
	return n.command
}

func (n *natureMock) InitialEnergy() int {
	return n.initialEnergy
}

type terrainMock struct {
}

func TestWithBrain(t *testing.T) {
	t.Parallel()

	brain := &brainMock{
		commands: []int{1, 2, 3},
	}
	agent := New(WithBrain(brain))

	if agent.brain != brain {
		t.Errorf("expected brain to be %v, got %v", brain, agent.brain)
	}
}

func TestIsAlive(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		energy  int
		isAlive bool
	}{
		{
			name:    "alive",
			energy:  1,
			isAlive: true,
		},
		{
			name:    "dead",
			energy:  0,
			isAlive: false,
		},
		{
			name:    "over dead",
			energy:  -1,
			isAlive: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agent := New()
			agent.energy = tt.energy

			if agent.IsAlive() != tt.isAlive {
				t.Errorf("expected IsAlive to be %v, got %v", tt.isAlive, agent.IsAlive())
			}
		})
	}
}

func TestEnergy(t *testing.T) {
	t.Parallel()

	agent := New()
	agent.energy = 1

	if agent.Energy() != 1 {
		t.Errorf("expected Energy to be %v, got %v", 1, agent.Energy())
	}
}

func TestIncreaseEnergy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		energy   int
		delta    int
		expected int
	}{
		{
			name:     "increase",
			energy:   1,
			delta:    1,
			expected: 2,
		},
		{
			name:     "over max",
			energy:   1,
			delta:    -2,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agent := New()
			agent.energy = tt.energy
			agent.IncreaseEnergy(tt.delta)

			if agent.energy != tt.expected {
				t.Errorf("expected energy to be %v, got %v", tt.expected, agent.energy)
			}
		})
	}
}

func TestDecreaseEnergy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		energy   int
		delta    int
		expected int
	}{
		{
			name:     "decrease",
			energy:   1,
			delta:    1,
			expected: 0,
		},
		{
			name:     "over min",
			energy:   1,
			delta:    -2,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agent := New()
			agent.energy = tt.energy
			agent.DecreaseEnergy(tt.delta)

			if agent.energy != tt.expected {
				t.Errorf("expected energy to be %v, got %v", tt.expected, agent.energy)
			}
		})
	}
}

func TestAction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                 string
		energy               int
		commands             []int
		maxDailyCommandCount int
		expectedActionsCount int
		wantErr              assert.ErrorAssertionFunc
	}{
		{
			name:                 "valid command count per day",
			energy:               10,
			commands:             []int{1, 2, 3},
			maxDailyCommandCount: 1,
			expectedActionsCount: 1,
			wantErr:              assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			brain := &brainMock{
				commands: tt.commands,
			}
			nature := &natureMock{
				maxDailyCommandCount: tt.maxDailyCommandCount,
				command: &commandMock{
					isInterrupt: false,
					delta:       1,
				},
			}
			agent := New(WithBrain(brain), WithNature(nature))
			agent.energy = tt.energy

			actionsCount, err := agent.Action(&terrainMock{})
			if (err != nil) == tt.wantErr(t, err) {
				return
			}

			assert.EqualValuesf(t, tt.expectedActionsCount, actionsCount, "expected actionsCount to be %v, got %v", tt.expectedActionsCount, actionsCount)
		})
	}
}
