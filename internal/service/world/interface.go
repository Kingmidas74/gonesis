package world

import (
	"errors"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
)

var ErrWorldIsNotInitialize = errors.New("can't create maze")

type Service interface {
	// Init initializes the world
	Init(config *configuration.Configuration, commands []contracts.Command) (contracts.World, error)
	// Update updates the world
	Update(config *configuration.Configuration) (contracts.World, error)
}
