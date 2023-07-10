package world

import (
	"encoding/json"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent/mutation"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent/reproduction"
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

func (s *srv) Init() (contracts.World, error) {
	mazeBuilder, err := s.getMazeBuilder()
	if err != nil {
		return nil, err
	}

	requiredEmptyCells := s.config.PlantConfiguration.InitialCount +
		s.config.HerbivoreConfiguration.InitialCount +
		s.config.CarnivoreConfiguration.InitialCount +
		s.config.DecomposerConfiguration.InitialCount +
		s.config.OmnivoreConfiguration.InitialCount

	m, err := mazeBuilder.SetWidth(s.config.WorldConfiguration.Ratio.Width).
		SetHeight(s.config.WorldConfiguration.Ratio.Height).
		SetRequiredEmptyCells(requiredEmptyCells).
		Build()
	if err != nil {
		return nil, err
	}

	terra, err := s.getTerrain(m)
	if err != nil {
		return nil, err
	}

	agents, err := s.generateAgents()
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

	s.world = world.New(terra, s.getAvailableCommands())
	return s.world, nil
}

func (s *srv) Update() (contracts.World, error) {
	if s.world == nil {
		c, _ := json.Marshal(*s.config)
		panic(string(c))
	}
	err := s.world.Next(s.config)
	return s.world, err
}

func (s *srv) UpdateConfiguration(config *configuration.Configuration) error {
	s.config = config
	return nil
}

func (s *srv) getMazeBuilder() (contracts.MazeBuilder, error) {
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

func (s *srv) generatePlants() ([]contracts.Agent, error) {
	reproductionSystem, err := s.getReproductionSystem(s.config.PlantConfiguration.ReproductionType)
	if err != nil {
		return nil, err
	}
	plants := make([]contracts.Agent, s.config.PlantConfiguration.InitialCount)
	plantNature := &nature.Plant{
		ReproductionSystem: reproductionSystem,
	}
	plantNature.Configure(s.config)
	for i := 0; i < s.config.PlantConfiguration.InitialCount; i++ {
		plants[i] = agent.NewAgent(plantNature)
	}
	return plants, nil
}

func (s *srv) generateHerbivores() ([]contracts.Agent, error) {
	reproductionSystem, err := s.getReproductionSystem(s.config.OmnivoreConfiguration.ReproductionType)
	if err != nil {
		return nil, err
	}
	herbivores := make([]contracts.Agent, s.config.HerbivoreConfiguration.InitialCount)
	herbivoreNature := &nature.Herbivore{
		ReproductionSystem: reproductionSystem,
	}
	herbivoreNature.Configure(s.config)
	for i := 0; i < s.config.HerbivoreConfiguration.InitialCount; i++ {
		herbivores[i] = agent.NewAgent(herbivoreNature)
	}
	return herbivores, nil
}

func (s *srv) generateCarnivores() ([]contracts.Agent, error) {
	reproductionSystem, err := s.getReproductionSystem(s.config.OmnivoreConfiguration.ReproductionType)
	if err != nil {
		return nil, err
	}
	carnivores := make([]contracts.Agent, s.config.CarnivoreConfiguration.InitialCount)
	carnivoreNature := &nature.Carnivore{
		ReproductionSystem: reproductionSystem,
	}
	carnivoreNature.Configure(s.config)
	for i := 0; i < s.config.CarnivoreConfiguration.InitialCount; i++ {
		carnivores[i] = agent.NewAgent(carnivoreNature)
	}
	return carnivores, nil
}

func (s *srv) generateDecomposers() ([]contracts.Agent, error) {
	reproductionSystem, err := s.getReproductionSystem(s.config.OmnivoreConfiguration.ReproductionType)
	if err != nil {
		return nil, err
	}
	decomposers := make([]contracts.Agent, s.config.DecomposerConfiguration.InitialCount)
	decomposerNature := &nature.Decomposer{
		ReproductionSystem: reproductionSystem,
	}
	decomposerNature.Configure(s.config)
	for i := 0; i < s.config.DecomposerConfiguration.InitialCount; i++ {
		decomposers[i] = agent.NewAgent(decomposerNature)
	}
	return decomposers, nil
}

func (s *srv) generateOmnivores() ([]contracts.Agent, error) {
	reproductionSystem, err := s.getReproductionSystem(s.config.OmnivoreConfiguration.ReproductionType)
	if err != nil {
		return nil, err
	}
	omnivores := make([]contracts.Agent, s.config.OmnivoreConfiguration.InitialCount)
	omnivoreNature := &nature.Omnivore{
		ReproductionSystem: reproductionSystem,
	}
	omnivoreNature.Configure(s.config)
	for i := 0; i < s.config.OmnivoreConfiguration.InitialCount; i++ {
		omnivores[i] = agent.NewAgent(omnivoreNature)
	}
	return omnivores, nil
}

func (s *srv) getReproductionSystem(reproductionSystemType enum.ReproductionSystemType) (contracts.ReproductionSystem, error) {
	switch reproductionSystemType {
	case enum.ReproductionSystemTypeBudding:
		return &reproduction.BuddingReproduction{
			Mutation: mutation.RandomizeMutation{},
		}, nil
	default:
		return nil, errors.ErrReproductionSystemTypeNotSupported

	}
}

func (s *srv) generateAgents() ([]contracts.Agent, error) {
	agents := make([]contracts.Agent, 0)

	plants, err := s.generatePlants()
	if err != nil {
		return agents, err
	}
	agents = append(agents, plants...)

	herbivores, err := s.generateHerbivores()
	if err != nil {
		return agents, err
	}
	agents = append(agents, herbivores...)

	carnivores, err := s.generateCarnivores()
	if err != nil {
		return agents, err
	}
	agents = append(agents, carnivores...)

	decomposers, err := s.generateDecomposers()
	if err != nil {
		return agents, err
	}
	agents = append(agents, decomposers...)

	omnivores, err := s.generateOmnivores()
	if err != nil {
		return agents, err
	}
	agents = append(agents, omnivores...)

	return agents, nil
}

func (s *srv) getAvailableCommands() []contracts.Command {
	return []contracts.Command{
		commands.NewPhotosynthesisCommand(),
		commands.NewEatCommand(),
		commands.NewMoveCommand(),
		commands.NewCallSubroutineCommand(),
		commands.NewEndSubroutineCommand(),
	}
}
