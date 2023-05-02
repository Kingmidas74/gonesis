package agent

import "testing"

type mockCommand struct {
	handleFunc      func(agent *Agent) int
	isInterruptFunc func() bool
}

func (c *mockCommand) Handle(agent *Agent) int {
	return c.handleFunc(agent)
}

func (c *mockCommand) IsInterrupt() bool {
	return c.isInterruptFunc()
}

func TestAgent_NextDay_UndefinedCommand(t *testing.T) {
	// Initialize test agent
	commands := []int{1, 2, 3}
	agent := New(10, commands)

	// Define findCommandPredicate function that always returns nil
	findCommandPredicate := func(identifier int) Command {
		return nil
	}

	// Test the case where the agent encounters an undefined command
	err := agent.NextDay(3, findCommandPredicate)

	// Check results
	if err != ErrCommandUndefined {
		t.Errorf("Expected error %v, but got %v", ErrCommandUndefined, err)
	}
}

func TestAgent_IsAlive(t *testing.T) {
	a := New(0, make([]int, 0))
	type fields struct {
		energy  int
		isAlive bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "greater zero",
			fields: fields{
				energy:  1,
				isAlive: true,
			},
		},
		{
			name: "equal zero",
			fields: fields{
				energy:  0,
				isAlive: false,
			},
		},
		{
			name: "less zero",
			fields: fields{
				energy:  -1,
				isAlive: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a.SetEnergy(tt.fields.energy)
			if got := a.IsAlive(); tt.fields.isAlive != got {
				t.Errorf("exp state %t, got state %t", tt.fields.isAlive, got)
			}
		})
	}
}
