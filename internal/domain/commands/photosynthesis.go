package commands

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

type PhotosynthesisCommand struct {
	isInterrupt bool
}

func NewPhotosynthesisCommand() *PhotosynthesisCommand {
	return &PhotosynthesisCommand{
		isInterrupt: true,
	}
}

func (c *PhotosynthesisCommand) Handle(agent contracts.Agent, terra contracts.Terrain) int {
	switch agent.AgentType() {
	case enum.AgentTypeHerbivore:
		return c.handleHerbivore(agent, terra)
	case enum.AgentTypeCarnivore:
		return c.handleCarnivore(agent, terra)
	case enum.AgentTypePlant:
		return c.handlePlant(agent, terra)
	}
	return 1
}

func (c *PhotosynthesisCommand) IsInterrupt() bool {
	return c.isInterrupt
}

func (c *PhotosynthesisCommand) handleHerbivore(agent contracts.Agent, terra contracts.Terrain) int {
	return 1
}

func (c *PhotosynthesisCommand) handleCarnivore(agent contracts.Agent, terra contracts.Terrain) int {
	return 1
}

func (c *PhotosynthesisCommand) handlePlant(agent contracts.Agent, terra contracts.Terrain) int {
	agent.IncreaseEnergy(terra.Cell(agent.X(), agent.Y()).Energy())
	return 1
}
