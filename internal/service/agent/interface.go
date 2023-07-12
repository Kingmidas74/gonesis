package agent

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
)

type Service interface {
	Generate(configuration *configuration.Configuration) ([]contracts.Agent, error)

	GeneratePlants(configuration *configuration.Configuration) ([]contracts.Agent, error)
	GenerateHerbivores(configuration *configuration.Configuration) ([]contracts.Agent, error)
	GenerateCarnivores(configuration *configuration.Configuration) ([]contracts.Agent, error)
	GenerateOmnivores(configuration *configuration.Configuration) ([]contracts.Agent, error)
	GenerateDecomposers(configuration *configuration.Configuration) ([]contracts.Agent, error)
}
