package world

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type agentMock struct {
	id           string
	isAlive      bool
	actionsCount int
}

func (a *agentMock) Action(terrain Terrain) (actionsCount int, err error) {
	return a.actionsCount, nil
}

func (a *agentMock) IsAlive() bool {
	return a.isAlive
}

func (a *agentMock) ID() string {
	return a.id
}

type cellMock struct {
	a Agent
}

func (c *cellMock) Agent() Agent {
	return c.a
}

func (c *cellMock) RemoveAgent() {
	c.a = nil
}

type terrainMock struct {
	cells []Cell
}

func (t *terrainMock) Cells() []Cell {
	return t.cells
}

func TestWorld_NextDay(t *testing.T) {
	tests := []struct {
		name    string
		terrain Terrain
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "ok",
			terrain: &terrainMock{
				cells: []Cell{
					&cellMock{
						a: &agentMock{
							id:           "test-agent-id",
							isAlive:      true,
							actionsCount: 1,
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := New(WithTerrain(tt.terrain))
			err := w.NextDay()
			if (err != nil) == tt.wantErr(t, err) {
				return
			}

			assert.EqualValuesf(t, 1, w.CurrentDay(), "expected currentDay to be %v, got %v", 1, w.CurrentDay())
		})
	}
}
