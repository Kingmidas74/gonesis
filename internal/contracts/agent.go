package contracts

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Agent interface {
	Coords
	Energy
	AgentNature
	Brain

	ID() string
	AgentType() enum.AgentType
	Generation() int

	NextDay(world Terrain, command func(commandIdentifier int) Command) error
	Kill(world Terrain)
	IsAlive() bool
	CreateChildren(world Terrain, config *configuration.Configuration) []Agent
}
