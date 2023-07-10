package world

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
)

type Service interface {
	// Init initializes the world
	Init() (contracts.World, error)
	// Update updates the world
	Update() (contracts.World, error)
	// UpdateConfiguration updates the configuration
	UpdateConfiguration(config *configuration.Configuration) error
}
