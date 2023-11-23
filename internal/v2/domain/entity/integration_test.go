package entity

import (
	"github.com/kingmidas74/gonesis-engine/internal/v2/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/v2/domain/entity/brain"
	"github.com/kingmidas74/gonesis-engine/internal/v2/domain/entity/nature"
	"testing"
)

func TestGame(t *testing.T) {
	t.Parallel()

	t.Run("New", func(t *testing.T) {
		natur := nature.NewBuilder().WithInitialEnergy(10).WithBrainVolume(2).Build()

		brain1 := brain.NewBuilder().WithCommands([]int{0, 1}).Build()

		agent1 := agent.New(agent.WithBrain(brain1), agent.WithNature(natur))
	})
}
