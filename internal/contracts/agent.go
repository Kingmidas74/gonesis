package contracts

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Agent interface {
	Coords
	Energy
	AgentNature

	Address() int
	Command(address *int) int
	Commands() []int
	AgentType() enum.AgentType

	NextDay(world Terrain, command func(commandIdentifier int) Command, config *configuration.Configuration) error
	Kill(world Terrain)
	IsAlive() bool
	CreateChildren(world Terrain, config *configuration.Configuration) []Agent
}
