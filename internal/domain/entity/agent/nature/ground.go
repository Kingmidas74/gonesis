package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type Ground struct {
}

func (a Ground) AgentType() enum.AgentType {
	return enum.AgentTypeGround
}

func (a Ground) Genesis(contracts.Agent) contracts.Agent {
	return nil
}
