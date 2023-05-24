package contracts

import "github.com/kingmidas74/gonesis-engine/internal/domain/enum"

type AgentNature interface {
	AgentType() enum.AgentType
	Genesis(a Agent) Agent
}
