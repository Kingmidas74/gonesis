package game

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/commands"
	"math/rand"

	"golang.org/x/exp/slices"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze/generator"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/terrain"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/terrain/topology"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/world"
)

const (
	initialEnergy = 2000
	brainVolume   = 20
)

func (s *Service) InitWorld(width, height int, agentsCount int) (contracts.World, error) {
	mazeBuilder := maze.NewMazeBuilder[generator.SidewinderGenerator]()
	m, err := mazeBuilder.SetWidth(width).
		SetHeight(height).
		SetRequiredEmptyCells(agentsCount).
		Build()
	if err != nil {
		return nil, err
	}

	terra := terrain.NewTerrain[topology.NeumannTopology](m)
	agents := s.generateAgents(agentsCount)

	emptyCells := terra.EmptyCells()
	pickedCellIndexes := make([]int, 0)
	for i := 0; i < agentsCount; i++ {
		targetIndex := rand.Intn(len(emptyCells))
		if slices.Contains(pickedCellIndexes, targetIndex) {
			i--
			continue
		}
		pickedCellIndexes = append(pickedCellIndexes, targetIndex)
		_ = emptyCells[targetIndex]
		agents[i].SetX(12)
		agents[i].SetY(22)
	}

	s.world = world.New(terra, agents, s.getAvailableCommands())
	return s.world, nil
}

func (s *Service) generateAgents(agentsCount int) []contracts.Agent {
	agents := make([]contracts.Agent, agentsCount)
	availableCommands := []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 3, 3, 3, 3, 2, 2, 2, 2}
	for i := 0; i < agentsCount; i++ {
		agents[i] = agent.New(initialEnergy, availableCommands, brainVolume)
	}
	return agents
}

func (s *Service) getAvailableCommands() []contracts.Command {
	return []contracts.Command{commands.NewMoveCommand()}
}
