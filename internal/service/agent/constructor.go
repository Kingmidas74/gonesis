package agent

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
)

type Service struct {
	agent contracts.Agent

	config *configuration.AgentConfiguration
}

func New(config *configuration.AgentConfiguration) *Service {
	return &Service{
		config: config,
	}
}
