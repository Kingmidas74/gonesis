package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Ground struct {
}

func (a Ground) AgentType() enum.AgentType {
	return enum.AgentTypeGround
}

func (a Ground) Genesis(contracts.Agent, *configuration.AgentConfiguration) []contracts.Agent {
	return make([]contracts.Agent, 0)
}

func (a Ground) MaxEnergy(*configuration.AgentConfiguration) int {
	return 0
}
