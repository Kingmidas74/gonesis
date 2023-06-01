package game

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"github.com/kingmidas74/gonesis-engine/internal/domain/errors"
	"math/rand"

	"golang.org/x/exp/slices"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/commands"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent/nature"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze/generator"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/terrain"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/terrain/topology"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/world"
)

func (s *Service) InitWorld(width, height int) (contracts.World, error) {
	mazeBuilder, err := s.getMazeBuilder()
	if err != nil {
		return nil, err
	}

	requiredEmptyCells := s.config.PlantConfiguration.InitialCount +
		s.config.HerbivoreConfiguration.InitialCount +
		s.config.CarnivoreConfiguration.InitialCount +
		s.config.DecomposerConfiguration.InitialCount +
		s.config.OmnivoreConfiguration.InitialCount

	m, err := mazeBuilder.SetWidth(width).
		SetHeight(height).
		SetRequiredEmptyCells(requiredEmptyCells).
		Build()
	if err != nil {
		return nil, err
	}

	terra, err := s.getTerrain(m)
	if err != nil {
		return nil, err
	}

	agents := s.generateAgents(requiredEmptyCells)

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
			panic("here")
		}
		emptyCell.SetAgent(agents[i])
	}

	s.world = world.New(terra, s.getAvailableCommands())
	return s.world, nil
}

func (s *Service) getMazeBuilder() (contracts.MazeBuilder, error) {
	switch s.config.WorldConfiguration.MazeType {
	case enum.MazeTypeBorder:
		return maze.NewMazeBuilder[generator.BorderGenerator](), nil
	case enum.MazeTypeBinary:
		return maze.NewMazeBuilder[generator.BinaryGenerator](), nil
	case enum.MazeTypeGrid:
		return maze.NewMazeBuilder[generator.GridGenerator](), nil
	case enum.MazeTypeAldousBroder:
		return maze.NewMazeBuilder[generator.AldousBroderGenerator](), nil
	case enum.MazeTypeSideWinder:
		return maze.NewMazeBuilder[generator.SidewinderGenerator](), nil
	case enum.MazeTypeEmpty:
		return maze.NewMazeBuilder[generator.EmptyGenerator](), nil
	default:
		return nil, errors.ErrMazeTypeNotSupported
	}
}
func (s *Service) getTerrain(m contracts.Maze) (contracts.Terrain, error) {
	switch s.config.WorldConfiguration.Topology {
	case enum.TopologyTypeMoore:
		return terrain.NewTerrain[topology.MooreTopology](m), nil
	case enum.TopologyTypeNeumann:
		return terrain.NewTerrain[topology.NeumannTopology](m), nil
	default:
		return nil, errors.ErrTopologyTypeNotSupported
	}
}

func (s *Service) generatePlants() []contracts.Agent {
	plants := make([]contracts.Agent, s.config.PlantConfiguration.InitialCount)
	plantNature := &nature.Plant{}
	plantNature.Configure(s.config)
	for i := 0; i < s.config.PlantConfiguration.InitialCount; i++ {
		plants[i] = agent.NewAgent(plantNature)
	}
	return plants
}

func (s *Service) generateHerbivores() []contracts.Agent {
	herbivores := make([]contracts.Agent, s.config.HerbivoreConfiguration.InitialCount)
	herbivoreNature := &nature.Herbivore{}
	herbivoreNature.Configure(s.config)
	for i := 0; i < s.config.HerbivoreConfiguration.InitialCount; i++ {
		herbivores[i] = agent.NewAgent(herbivoreNature)
	}
	return herbivores
}

func (s *Service) generateCarnivores() []contracts.Agent {
	carnivores := make([]contracts.Agent, s.config.CarnivoreConfiguration.InitialCount)
	carnivoreNature := &nature.Carnivore{}
	carnivoreNature.Configure(s.config)
	for i := 0; i < s.config.CarnivoreConfiguration.InitialCount; i++ {
		carnivores[i] = agent.NewAgent(carnivoreNature)
	}
	return carnivores
}

func (s *Service) generateDecomposers() []contracts.Agent {
	decomposers := make([]contracts.Agent, s.config.DecomposerConfiguration.InitialCount)
	decomposerNature := &nature.Decomposer{}
	decomposerNature.Configure(s.config)
	for i := 0; i < s.config.DecomposerConfiguration.InitialCount; i++ {
		decomposers[i] = agent.NewAgent(decomposerNature)
	}
	return decomposers
}

func (s *Service) generateOmnivores() []contracts.Agent {
	omnivores := make([]contracts.Agent, s.config.OmnivoreConfiguration.InitialCount)
	omnivoreNature := &nature.Omnivore{}
	omnivoreNature.Configure(s.config)
	for i := 0; i < s.config.OmnivoreConfiguration.InitialCount; i++ {
		omnivores[i] = agent.NewAgent(omnivoreNature)
	}
	return omnivores
}

func (s *Service) generateAgents(agentsCount int) []contracts.Agent {
	agents := make([]contracts.Agent, 0, agentsCount)

	agents = append(agents, s.generatePlants()...)
	agents = append(agents, s.generateHerbivores()...)
	agents = append(agents, s.generateCarnivores()...)
	agents = append(agents, s.generateDecomposers()...)
	agents = append(agents, s.generateOmnivores()...)

	return agents
}

func (s *Service) getAvailableCommands() []contracts.Command {
	return []contracts.Command{
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
	}
}
