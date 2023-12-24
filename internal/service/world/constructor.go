package world

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/contract"
	"github.com/kingmidas74/gonesis-engine/internal/service/agent"
)

type Params struct {
	AgentService agent.Service

	MazeGeneratorCollection MazeGeneratorCollection
}

type srv struct {
	world contract.World

	agentService            agent.Service
	mazeGeneratorCollection MazeGeneratorCollection
}

func New(params Params) Service {
	return &srv{
		agentService:            params.AgentService,
		mazeGeneratorCollection: params.MazeGeneratorCollection,
	}
}

type MazeGeneratorCollection interface {
	Get(mazeType contract.MazeType) (contract.MazeGenerator, error)
}
