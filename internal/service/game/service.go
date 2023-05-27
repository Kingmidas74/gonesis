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
	m, err := mazeBuilder.SetWidth(width).
		SetHeight(height).
		SetRequiredEmptyCells(s.config.AgentConfiguration.InitialCount).
		Build()
	if err != nil {
		return nil, err
	}

	terra := terrain.NewTerrain[topology.MooreTopology](m)
	agents := s.generateAgents(s.config.AgentConfiguration.InitialCount)

	emptyCells := terra.EmptyCells()
	pickedCellIndexes := make([]int, 0)
	for i := 0; i < s.config.AgentConfiguration.InitialCount; i++ {
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

func (s *Service) generateAgents(agentsCount int) []contracts.Agent {
	agents := make([]contracts.Agent, agentsCount)
	for i := 0; i < agentsCount; i++ {
		switch rand.Intn(18) + 1 {
		case 1, 2, 3, 4, 5:
			agents[i] = agent.NewAgent[nature.Herbivore](s.config.AgentConfiguration.InitialEnergy, s.config.AgentConfiguration.BrainVolume)
			continue
		case 6, 7, 8, 9, 10:
			agents[i] = agent.NewAgent[nature.Carnivore](s.config.AgentConfiguration.InitialEnergy, s.config.AgentConfiguration.BrainVolume)
			continue
		case 0:
			agents[i] = agent.NewAgent[nature.Decomposer](s.config.AgentConfiguration.InitialEnergy, s.config.AgentConfiguration.BrainVolume)
			continue
		case 11, 12, 13, 14:
			agents[i] = agent.NewAgent[nature.Plant](s.config.AgentConfiguration.InitialEnergy/2, s.config.AgentConfiguration.BrainVolume)
			continue
		case 16, 17, 18, 15:
			agents[i] = agent.NewAgent[nature.Omnivore](s.config.AgentConfiguration.InitialEnergy, s.config.AgentConfiguration.BrainVolume)
			continue
		}

	}
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
