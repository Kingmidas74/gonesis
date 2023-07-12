package world

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"github.com/kingmidas74/gonesis-engine/internal/domain/errors"
	"math/rand"

	"golang.org/x/exp/slices"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/terrain"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/terrain/topology"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/world"
)

func (s *srv) Init(config *configuration.Configuration, availableCommands []contracts.Command) (contracts.World, error) {
	requiredEmptyCells := config.PlantConfiguration.InitialCount +
		config.HerbivoreConfiguration.InitialCount +
		config.CarnivoreConfiguration.InitialCount +
		config.DecomposerConfiguration.InitialCount +
		config.OmnivoreConfiguration.InitialCount

	m, err := s.mazeService.Generate(config.WorldConfiguration.MazeType, config.WorldConfiguration.Ratio.Width, config.WorldConfiguration.Ratio.Height, requiredEmptyCells)
	if err != nil {
		return nil, err
	}

	terra, err := s.getTerrain(m)
	if err != nil {
		return nil, err
	}

	agents, err := s.agentService.Generate(config)
	if err != nil {
		return nil, err
	}

	emptyCells := terra.EmptyCells()
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

	s.world = world.New(terra, availableCommands)
	return s.world, nil
}

func (s *srv) Update(config *configuration.Configuration) (contracts.World, error) {
	if s.world == nil {
		return nil, ErrWorldIsNotInitialize
	}

	err := s.world.Next(config)
	return s.world, err
}

func (s *srv) getTerrain(m contracts.Maze) (contracts.Terrain, error) {
	switch s.config.WorldConfiguration.Topology {
	case enum.TopologyTypeMoore:
		return terrain.NewTerrain[topology.MooreTopology](m), nil
	case enum.TopologyTypeNeumann:
		return terrain.NewTerrain[topology.NeumannTopology](m), nil
	default:
		return nil, errors.ErrTopologyTypeNotSupported
	}
}
