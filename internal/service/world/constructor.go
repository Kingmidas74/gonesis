package world

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/service/agent"
	"github.com/kingmidas74/gonesis-engine/internal/service/maze"
)

type srv struct {
	world contracts.World

	mazeService  maze.Service
	agentService agent.Service

	config *configuration.Configuration
}

func New(config *configuration.Configuration) Service {
	return &srv{
		mazeService:  maze.New(config),
		agentService: agent.New(),

		config: config,
	}
}
