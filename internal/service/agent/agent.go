package agent

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	contract "github.com/kingmidas74/gonesis-engine/internal/domain/contract"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/brain"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/command/eat"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/command/move"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/command/photosynthesis"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/mutation/randomize"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/nature"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/reproduction/budding"
)

func (s *srv) Generate(config *configuration.Configuration) ([]contract.Agent, error) {
	agents := make([]contract.Agent, 0)

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

func (s *srv) GeneratePlants(config *configuration.Configuration) ([]contract.Agent, error) {
	availableCommands := []contract.Command{
		photosynthesis.New(true),
	}
	commandSequence := s.randomIntGenerator.GenerateRandomIntSequence(config.PlantConfiguration.BrainVolume, []int{0})

	mutationHandler := randomize.New(s.randomIntGenerator)
	reproductionSystem := budding.New(s.randomFloatGenerator, mutationHandler)

	n, err := nature.New(reproductionSystem,
		nature.WithBrainVolume(config.PlantConfiguration.BrainVolume),
		nature.WithCommands(availableCommands),
		nature.WithInitialEnergy(config.PlantConfiguration.InitialEnergy),
		nature.WithMaxDailyCommandCount(config.PlantConfiguration.MaxDailyCommandCount),
		nature.WithMutationChance(config.PlantConfiguration.MutationChance),
		nature.WithReproductionChance(config.PlantConfiguration.ReproductionChance),
		nature.WithType(contract.AgentTypePlant),
		nature.WithInitialCount(config.PlantConfiguration.InitialCount),
		nature.WithReproductionEnergyCost(config.PlantConfiguration.ReproductionEnergyCost))
	if err != nil {
		return nil, err
	}

	agents := make([]contract.Agent, config.PlantConfiguration.InitialCount)
	for i := 0; i < config.PlantConfiguration.InitialCount; i++ {
		agents[i] = agent.New(
			agent.WithBrain(
				brain.New(
					brain.WithCommands(commandSequence),
				),
			),
			agent.WithNature(n),
		)
	}
	return agents, nil
}

func (s *srv) GenerateHerbivores(config *configuration.Configuration) ([]contract.Agent, error) {
	availableCommands := []contract.Command{
		move.New(true),
		eat.New(true),
		move.New(true),
		eat.New(true),
		move.New(true),
		eat.New(true),
		eat.New(true),
		eat.New(true),
		eat.New(true),
		eat.New(true),
	}
	commandSequence := s.randomIntGenerator.GenerateRandomIntSequence(config.HerbivoreConfiguration.BrainVolume, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	mutationHandler := randomize.New(s.randomIntGenerator)
	reproductionSystem := budding.New(s.randomFloatGenerator, mutationHandler)

	n, err := nature.New(reproductionSystem,
		nature.WithBrainVolume(config.HerbivoreConfiguration.BrainVolume),
		nature.WithCommands(availableCommands),
		nature.WithInitialEnergy(config.HerbivoreConfiguration.InitialEnergy),
		nature.WithMaxDailyCommandCount(config.HerbivoreConfiguration.MaxDailyCommandCount),
		nature.WithMutationChance(config.HerbivoreConfiguration.MutationChance),
		nature.WithReproductionChance(config.HerbivoreConfiguration.ReproductionChance),
		nature.WithType(contract.AgentTypeHerbivore),
		nature.WithInitialCount(config.HerbivoreConfiguration.InitialCount),
		nature.WithAvailableFood(map[contract.AgentType]int{
			contract.AgentTypePlant: 2,
		}),
		nature.WithReproductionEnergyCost(config.HerbivoreConfiguration.ReproductionEnergyCost))
	if err != nil {
		return nil, err
	}

	agents := make([]contract.Agent, config.HerbivoreConfiguration.InitialCount)
	for i := 0; i < config.HerbivoreConfiguration.InitialCount; i++ {
		agents[i] = agent.New(
			agent.WithBrain(
				brain.New(
					brain.WithCommands(commandSequence),
				),
			),
			agent.WithNature(n),
		)
	}
	return agents, nil
}

func (s *srv) GenerateCarnivores(config *configuration.Configuration) ([]contract.Agent, error) {
	availableCommands := []contract.Command{
		move.New(true),
		eat.New(true),
		move.New(true),
		eat.New(true),
		move.New(true),
		eat.New(true),
		eat.New(true),
		eat.New(true),
		eat.New(true),
		eat.New(true),
	}
	commandSequence := s.randomIntGenerator.GenerateRandomIntSequence(config.CarnivoreConfiguration.BrainVolume, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	mutationHandler := randomize.New(s.randomIntGenerator)
	reproductionSystem := budding.New(s.randomFloatGenerator, mutationHandler)

	n, err := nature.New(reproductionSystem,
		nature.WithBrainVolume(config.CarnivoreConfiguration.BrainVolume),
		nature.WithCommands(availableCommands),
		nature.WithInitialEnergy(config.CarnivoreConfiguration.InitialEnergy),
		nature.WithMaxDailyCommandCount(config.CarnivoreConfiguration.MaxDailyCommandCount),
		nature.WithMutationChance(config.CarnivoreConfiguration.MutationChance),
		nature.WithReproductionChance(config.CarnivoreConfiguration.ReproductionChance),
		nature.WithType(contract.AgentTypeCarnivore),
		nature.WithInitialCount(config.CarnivoreConfiguration.InitialCount),
		nature.WithAvailableFood(map[contract.AgentType]int{
			contract.AgentTypeHerbivore: 2,
			contract.AgentTypeCarnivore: 1,
		}),
		nature.WithReproductionEnergyCost(config.CarnivoreConfiguration.ReproductionEnergyCost))
	if err != nil {
		return nil, err
	}

	agents := make([]contract.Agent, config.CarnivoreConfiguration.InitialCount)
	for i := 0; i < config.CarnivoreConfiguration.InitialCount; i++ {
		agents[i] = agent.New(
			agent.WithBrain(
				brain.New(
					brain.WithCommands(commandSequence),
				),
			),
			agent.WithNature(n),
		)
	}
	return agents, nil
}

func (s *srv) GenerateDecomposers(config *configuration.Configuration) ([]contract.Agent, error) {
	availableCommands := []contract.Command{
		move.New(true),
		eat.New(true),
		move.New(true),
		eat.New(true),
		move.New(true),
		eat.New(true),
		eat.New(true),
		eat.New(true),
		eat.New(true),
		eat.New(true),
	}
	commandSequence := s.randomIntGenerator.GenerateRandomIntSequence(config.DecomposerConfiguration.BrainVolume, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	mutationHandler := randomize.New(s.randomIntGenerator)
	reproductionSystem := budding.New(s.randomFloatGenerator, mutationHandler)

	n, err := nature.New(reproductionSystem,
		nature.WithBrainVolume(config.DecomposerConfiguration.BrainVolume),
		nature.WithCommands(availableCommands),
		nature.WithInitialEnergy(config.DecomposerConfiguration.InitialEnergy),
		nature.WithMaxDailyCommandCount(config.DecomposerConfiguration.MaxDailyCommandCount),
		nature.WithMutationChance(config.DecomposerConfiguration.MutationChance),
		nature.WithReproductionChance(config.DecomposerConfiguration.ReproductionChance),
		nature.WithType(contract.AgentTypeDecomposer),
		nature.WithInitialCount(config.DecomposerConfiguration.InitialCount),
		nature.WithReproductionEnergyCost(config.DecomposerConfiguration.ReproductionEnergyCost))
	if err != nil {
		return nil, err
	}

	agents := make([]contract.Agent, config.DecomposerConfiguration.InitialCount)
	for i := 0; i < config.DecomposerConfiguration.InitialCount; i++ {
		agents[i] = agent.New(
			agent.WithBrain(
				brain.New(
					brain.WithCommands(commandSequence),
				),
			),
			agent.WithNature(n),
		)
	}
	return agents, nil
}

func (s *srv) GenerateOmnivores(config *configuration.Configuration) ([]contract.Agent, error) {
	availableCommands := []contract.Command{
		move.New(true),
		eat.New(true),
		move.New(true),
		eat.New(true),
		move.New(true),
		eat.New(true),
		eat.New(true),
		eat.New(true),
		eat.New(true),
		eat.New(true),
	}
	commandSequence := s.randomIntGenerator.GenerateRandomIntSequence(config.OmnivoreConfiguration.BrainVolume, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	mutationHandler := randomize.New(s.randomIntGenerator)
	reproductionSystem := budding.New(s.randomFloatGenerator, mutationHandler)

	n, err := nature.New(reproductionSystem,
		nature.WithBrainVolume(config.OmnivoreConfiguration.BrainVolume),
		nature.WithCommands(availableCommands),
		nature.WithInitialEnergy(config.OmnivoreConfiguration.InitialEnergy),
		nature.WithMaxDailyCommandCount(config.OmnivoreConfiguration.MaxDailyCommandCount),
		nature.WithMutationChance(config.OmnivoreConfiguration.MutationChance),
		nature.WithReproductionChance(config.OmnivoreConfiguration.ReproductionChance),
		nature.WithType(contract.AgentTypeOmnivore),
		nature.WithInitialCount(config.OmnivoreConfiguration.InitialCount),
		nature.WithAvailableFood(map[contract.AgentType]int{
			contract.AgentTypePlant:     1,
			contract.AgentTypeHerbivore: 2,
			contract.AgentTypeCarnivore: 2,
			contract.AgentTypeOmnivore:  1,
		}),
		nature.WithReproductionEnergyCost(config.OmnivoreConfiguration.ReproductionEnergyCost))
	if err != nil {
		return nil, err
	}

	agents := make([]contract.Agent, config.OmnivoreConfiguration.InitialCount)
	for i := 0; i < config.OmnivoreConfiguration.InitialCount; i++ {
		agents[i] = agent.New(
			agent.WithBrain(
				brain.New(
					brain.WithCommands(commandSequence),
				),
			),
			agent.WithNature(n),
		)
	}
	return agents, nil
}
