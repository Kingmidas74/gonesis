package game

import (
	"errors"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
)

var ErrNotEnoughEmptyCells = errors.New("not enough empty cells")

type Service struct {
	world contracts.World
}

func New() *Service {
	return &Service{}
}
