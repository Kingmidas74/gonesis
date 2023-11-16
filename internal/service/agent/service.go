package agent

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/commands"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent/mutation"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent/nature"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent/reproduction"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"github.com/kingmidas74/gonesis-engine/internal/domain/errors"
)

func (s *srv) Generate(config *configuration.Configuration) ([]contracts.Agent, error) {
	agents := make([]contracts.Agent, 0)

	plants, err := s.GeneratePlants(config)
	if err != nil {
		return agents, err
	}
	agents = append(agents, plants...)

	herbivores, err := s.GenerateHerbivores(config)
	if err != nil {
		return agents, err
	}
	agents = append(agents, herbivores...)

	carnivores, err := s.GenerateCarnivores(config)
	if err != nil {
		return agents, err
	}
	agents = append(agents, carnivores...)

	decomposers, err := s.GenerateDecomposers(config)
	if err != nil {
		return agents, err
	}
	agents = append(agents, decomposers...)

	omnivores, err := s.GenerateOmnivores(config)
	if err != nil {
		return agents, err
	}
	agents = append(agents, omnivores...)

	return agents, nil
}

func (s *srv) GeneratePlants(config *configuration.Configuration) ([]contracts.Agent, error) {
	availableCommands := []contracts.Command{
		commands.NewPhotosynthesisCommand(),
		commands.NewPhotosynthesisCommand(),
		commands.NewPhotosynthesisCommand(),
		commands.NewPhotosynthesisCommand(),
	}

	reproductionSystem, err := s.getReproductionSystem(config.PlantConfiguration.ReproductionType)
	if err != nil {
		return nil, err
	}
	plants := make([]contracts.Agent, config.PlantConfiguration.InitialCount)
	plantNature := &nature.Plant{
		ReproductionSystem: reproductionSystem,
	}

	if err := plantNature.Configure(config, availableCommands); err != nil {
		return nil, err
	}

	for i := 0; i < config.PlantConfiguration.InitialCount; i++ {
		plants[i] = agent.NewAgent(plantNature)
	}
	return plants, nil
}

func (s *srv) GenerateHerbivores(config *configuration.Configuration) ([]contracts.Agent, error) {
	availableCommands := []contracts.Command{
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
	}

	reproductionSystem, err := s.getReproductionSystem(config.OmnivoreConfiguration.ReproductionType)
	if err != nil {
		return nil, err
	}
	herbivores := make([]contracts.Agent, config.HerbivoreConfiguration.InitialCount)
	herbivoreNature := &nature.Herbivore{
		ReproductionSystem: reproductionSystem,
	}

	if err := herbivoreNature.Configure(config, availableCommands); err != nil {
		return nil, err
	}

	for i := 0; i < config.HerbivoreConfiguration.InitialCount; i++ {
		herbivores[i] = agent.NewAgent(herbivoreNature)
	}
	return herbivores, nil
}

func (s *srv) GenerateCarnivores(config *configuration.Configuration) ([]contracts.Agent, error) {
	availableCommands := []contracts.Command{
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
	}

	reproductionSystem, err := s.getReproductionSystem(config.OmnivoreConfiguration.ReproductionType)
	if err != nil {
		return nil, err
	}

	carnivores := make([]contracts.Agent, config.CarnivoreConfiguration.InitialCount)
	carnivoreNature := &nature.Carnivore{
		ReproductionSystem: reproductionSystem,
	}

	if err := carnivoreNature.Configure(config, availableCommands); err != nil {
		return nil, err
	}

	for i := 0; i < config.CarnivoreConfiguration.InitialCount; i++ {
		carnivores[i] = agent.NewAgent(carnivoreNature)
	}
	return carnivores, nil
}

func (s *srv) GenerateDecomposers(config *configuration.Configuration) ([]contracts.Agent, error) {
	var availableCommands []contracts.Command

	reproductionSystem, err := s.getReproductionSystem(config.OmnivoreConfiguration.ReproductionType)
	if err != nil {
		return nil, err
	}
	decomposers := make([]contracts.Agent, config.DecomposerConfiguration.InitialCount)
	decomposerNature := &nature.Decomposer{
		ReproductionSystem: reproductionSystem,
	}

	if err := decomposerNature.Configure(config, availableCommands); err != nil {
		return nil, err
	}

	for i := 0; i < config.DecomposerConfiguration.InitialCount; i++ {
		decomposers[i] = agent.NewAgent(decomposerNature)
	}
	return decomposers, nil
}

func (s *srv) GenerateOmnivores(config *configuration.Configuration) ([]contracts.Agent, error) {
	availableCommands := []contracts.Command{
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
		commands.NewEatCommand(),
	}

	reproductionSystem, err := s.getReproductionSystem(config.OmnivoreConfiguration.ReproductionType)
	if err != nil {
		return nil, err
	}
	omnivores := make([]contracts.Agent, config.OmnivoreConfiguration.InitialCount)
	omnivoreNature := &nature.Omnivore{
		ReproductionSystem: reproductionSystem,
	}

	if err := omnivoreNature.Configure(config, availableCommands); err != nil {
		return nil, err
	}

	for i := 0; i < config.OmnivoreConfiguration.InitialCount; i++ {
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
