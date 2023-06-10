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

func (s *Service) CreateAgent() (contracts.Agent, error) {
	reproductionSystem, err := s.getReproductionSystem(s.config.ReproductionType)
	if err != nil {
		return nil, err
	}
	omnivoreNature := &nature.Omnivore{
		ReproductionSystem: reproductionSystem,
	}
	c := configuration.NewConfiguration()
	c.OmnivoreConfiguration = *s.config
	omnivoreNature.Configure(c)

	s.agent = agent.NewAgent(omnivoreNature)
	return s.agent, nil
}

func (s *Service) getReproductionSystem(reproductionSystemType enum.ReproductionSystemType) (contracts.ReproductionSystem, error) {
	switch reproductionSystemType {
	case enum.ReproductionSystemTypeBudding:
		return &reproduction.BuddingReproduction{
			Mutation: mutation.RandomizeMutation{},
		}, nil
	default:
		return nil, errors.ErrReproductionSystemTypeNotSupported
	}
}

func (s *Service) getAvailableCommands() []contracts.Command {
	return []contracts.Command{
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewMoveCommand(),
		commands.NewEndSubroutineCommand(),
		commands.NewEatCommand(),
		commands.NewCallSubroutineCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewCallSubroutineCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewEndSubroutineCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewCallSubroutineCommand(),
		commands.NewMoveCommand(),
		commands.NewEatCommand(),
		commands.NewMoveCommand(),
		commands.NewEndSubroutineCommand(),
		commands.NewEatCommand(),
		commands.NewCallSubroutineCommand(),
	}
}
