package game

import (
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
	mazeBuilder := maze.NewMazeBuilder[generator.BorderGenerator]()
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

	terra := terrain.NewTerrain[topology.MooreTopology](m)
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

func (s *Service) generatePlants() []contracts.Agent {
	plants := make([]contracts.Agent, s.config.PlantConfiguration.InitialCount)
	for i := 0; i < s.config.PlantConfiguration.InitialCount; i++ {
		plants[i] = agent.NewAgent[nature.Plant](s.config.PlantConfiguration.InitialEnergy/2, s.config.PlantConfiguration.BrainVolume)
	}
	return plants
}

func (s *Service) generateHerbivores() []contracts.Agent {
	herbivores := make([]contracts.Agent, s.config.HerbivoreConfiguration.InitialCount)
	for i := 0; i < s.config.HerbivoreConfiguration.InitialCount; i++ {
		herbivores[i] = agent.NewAgent[nature.Herbivore](s.config.HerbivoreConfiguration.InitialEnergy, s.config.HerbivoreConfiguration.BrainVolume)
	}
	return herbivores
}

func (s *Service) generateCarnivores() []contracts.Agent {
	carnivores := make([]contracts.Agent, s.config.CarnivoreConfiguration.InitialCount)
	for i := 0; i < s.config.CarnivoreConfiguration.InitialCount; i++ {
		carnivores[i] = agent.NewAgent[nature.Carnivore](s.config.CarnivoreConfiguration.InitialEnergy, s.config.CarnivoreConfiguration.BrainVolume)
	}
	return carnivores
}

func (s *Service) generateDecomposers() []contracts.Agent {
	decomposers := make([]contracts.Agent, s.config.DecomposerConfiguration.InitialCount)
	for i := 0; i < s.config.DecomposerConfiguration.InitialCount; i++ {
		decomposers[i] = agent.NewAgent[nature.Decomposer](s.config.DecomposerConfiguration.InitialEnergy, s.config.DecomposerConfiguration.BrainVolume)
	}
	return decomposers
}

func (s *Service) generateOmnivores() []contracts.Agent {
	omnivores := make([]contracts.Agent, s.config.OmnivoreConfiguration.InitialCount)
	for i := 0; i < s.config.OmnivoreConfiguration.InitialCount; i++ {
		omnivores[i] = agent.NewAgent[nature.Omnivore](s.config.OmnivoreConfiguration.InitialEnergy, s.config.OmnivoreConfiguration.BrainVolume)
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
