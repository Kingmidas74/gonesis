package maze

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
)

type Service[K contracts.MazeGenerator] struct{}

func NewMazeService[K contracts.MazeGenerator]() *Service[K] {
	return &Service[K]{}
}
