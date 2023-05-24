package contracts

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type AgentNature interface {
	AgentType() enum.AgentType
	Genesis(a Agent, config *configuration.AgentConfiguration) []Agent
}
