package agent

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/contract"
)

type Service interface {
	Generate(configuration *configuration.Configuration) ([]contract.Agent, error)

	GeneratePlants(configuration *configuration.Configuration) ([]contract.Agent, error)
	GenerateHerbivores(configuration *configuration.Configuration) ([]contract.Agent, error)
	GenerateCarnivores(configuration *configuration.Configuration) ([]contract.Agent, error)
	GenerateOmnivores(configuration *configuration.Configuration) ([]contract.Agent, error)
	GenerateDecomposers(configuration *configuration.Configuration) ([]contract.Agent, error)
}
