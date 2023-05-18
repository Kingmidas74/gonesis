package agent

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/commands"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"testing"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
)

type mockTerrain struct {
	GetNeighborFunc func(x, y int, direction int) contracts.Cell
}

func (m mockTerrain) GetNeighbors(x, y int) []contracts.Cell {
	//TODO implement me
	panic("implement me")
}

func (m mockTerrain) Cell(x, y int) contracts.Cell {
	//TODO implement me
	panic("implement me")
}

func (m mockTerrain) Cells() []contracts.Cell {
	//TODO implement me
	panic("implement me")
}

func (m mockTerrain) Width() int {
	//TODO implement me
	panic("implement me")
}

func (m mockTerrain) Height() int {
	//TODO implement me
	panic("implement me")
}

func (m mockTerrain) SetCellType(x, y int, cell enum.CellType) {
	//TODO implement me
	panic("implement me")
}

func (m mockTerrain) GetNeighbor(x, y int, direction int) contracts.Cell {
	return m.GetNeighborFunc(x, y, direction)
}

func (m mockTerrain) EmptyCells() []contracts.Cell {
	return make([]contracts.Cell, 0)
}

func TestAgent_NextDay_UndefinedCommand(t *testing.T) {
	// Initialize test agent
	agent := New(10, 20)

	// Define findCommandPredicate function that always returns nil
	findCommandPredicate := func(identifier int) contracts.Command {
		return nil
	}

	// Test the case where the agent encounters an undefined command
	err := agent.NextDay(3, mockTerrain{}, findCommandPredicate)

	// Check results
	if err != ErrCommandUndefined {
		t.Errorf("Expected error %v, but got %v", ErrCommandUndefined, err)
	}
}

func TestAgent_IsAlive(t *testing.T) {
	a := New(0, 20)
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

func TestAgent_BrainCorrection(t *testing.T) {
	a := New(10, 20)
	for {
		err := a.NextDay(11, mockTerrain{
			GetNeighborFunc: func(x, y int, direction int) contracts.Cell {
				return entity.NewCell(a.X(), a.Y()-1, enum.CellTypeEmpty)
			},
		}, func(identifier int) contracts.Command {
			if identifier == 0 {
				return commands.NewMoveCommand()
			}
			return nil
		})
		if err != nil {
			t.Errorf("there should be no error, got %v", err)
		}
		if !a.IsAlive() {
			break
		}
	}
	if a.Energy() != 0 {
		t.Errorf("exp energy 0, got %d", a.Energy())
	}
}
