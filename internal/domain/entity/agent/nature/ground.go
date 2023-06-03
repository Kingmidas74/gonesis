package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Ground struct {
	contracts.ReproductionSystem

	config *configuration.Configuration
}

func (a *Ground) Configure(config *configuration.Configuration) {
	a.config = config
}

func (a *Ground) AgentType() enum.AgentType {
	return enum.AgentTypeGround
}

func (a *Ground) MaxEnergy() int {
	return 0
}

func (a *Ground) BrainVolume() int {
	return 0
}

func (a *Ground) MaxDailyCommandCount() int {
	return 0
}

func (a *Ground) InitialEnergy() int {
	return 0
}
