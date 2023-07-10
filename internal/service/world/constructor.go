package world

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
)

type srv struct {
	world contracts.World

	config *configuration.Configuration
}

func New(config *configuration.Configuration) Service {
	return &srv{
		config: config,
	}
}
