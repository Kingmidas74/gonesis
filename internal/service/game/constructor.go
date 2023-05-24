package game

import (
	"errors"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
)

var ErrNotEnoughEmptyCells = errors.New("not enough empty cells")

type Service struct {
	world contracts.World

	config *configuration.Configuration
}

func New(config *configuration.Configuration) *Service {
	return &Service{
		config: config,
	}
}
