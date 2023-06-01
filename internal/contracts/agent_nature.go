package contracts

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Ptr[T any] interface {
	*T
}

type AgentNature interface {
	Configure(config *configuration.Configuration)

	AgentType() enum.AgentType
	Genesis(a Agent) []Agent
	MaxEnergy() int
	MaxDailyCommandCount() int
	InitialEnergy() int
	BrainVolume() int
}
