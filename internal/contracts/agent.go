package contracts

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Agent interface {
	Coords
	Energy

	Address() int
	Command(address *int) int
	Commands() []int
	AgentType() enum.AgentType

	NextDay(maxSteps int, world Terrain, command func(commandIdentifier int) Command) error
	Kill(world Terrain)
	IsAlive() bool
	CreateChild(world Terrain) Agent
}
