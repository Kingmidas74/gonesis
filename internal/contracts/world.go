package contracts

import "github.com/kingmidas74/gonesis-engine/internal/domain/configuration"

type World interface {
	Terrain

	Agents() []Agent
	Next(config *configuration.Configuration) error
	CurrentDay() int
}
