package world

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/contract"
	"github.com/kingmidas74/gonesis-engine/internal/service/agent"
)

type Params struct {
	AgentService agent.Service
}

type srv struct {
	world contract.World

	agentService agent.Service
}

func New(params Params) Service {
	return &srv{
		agentService: params.AgentService,
	}
}
