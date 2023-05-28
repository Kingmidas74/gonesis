package commands

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"golang.org/x/exp/slices"
)

const EatEnergy = 100

type EatCommand struct {
	isInterrupt bool
}

func NewEatCommand() *EatCommand {
	return &EatCommand{
		isInterrupt: true,
	}
}

func (c *EatCommand) Handle(agent contracts.Agent, terra contracts.Terrain) int {
	switch agent.AgentType() {
	case enum.AgentTypeHerbivore:
		return c.handleHerbivore(agent, terra)
	case enum.AgentTypeCarnivore:
		return c.handleCarnivore(agent, terra)
	case enum.AgentTypePlant:
		return c.handlePlant(agent, terra)
	case enum.AgentTypeOmnivore:
		return c.handleOmnivore(agent, terra)
	}
	return 1
}

func (c *EatCommand) IsInterrupt() bool {
	return c.isInterrupt
}

func (c *EatCommand) handleHerbivore(agent contracts.Agent, terra contracts.Terrain) int {
	whereAddress := agent.Address() + 1
	direction := agent.Command(&whereAddress)
	targetCell := terra.GetNeighbor(agent.X(), agent.Y(), direction)
	if targetCell == nil {
		return 0
	}

	if !targetCell.IsAgent() {
		return 1
	}

	if targetCell.Agent().AgentType() == enum.AgentTypePlant {
		agent.IncreaseEnergy(targetCell.Agent().Energy())
		targetCell.Agent().Kill(terra)
		terra.Cell(agent.X(), agent.Y()).RemoveAgent()
		targetCell.SetAgent(agent)
	}
	return 1
}

func (c *EatCommand) handleCarnivore(agent contracts.Agent, terra contracts.Terrain) int {
	whereAddress := agent.Address() + 1
	direction := agent.Command(&whereAddress)
	targetCell := terra.GetNeighbor(agent.X(), agent.Y(), direction)
	if targetCell == nil {
		return 0
	}

	if !targetCell.IsAgent() {
		return 1
	}

	if targetCell.Agent().AgentType() == enum.AgentTypeHerbivore {
		agent.IncreaseEnergy(targetCell.Agent().Energy())
		targetCell.Agent().Kill(terra)
		terra.Cell(agent.X(), agent.Y()).RemoveAgent()
		targetCell.SetAgent(agent)
	}
	return 1
}

func (c *EatCommand) handlePlant(agent contracts.Agent, terra contracts.Terrain) int {
	agent.IncreaseEnergy(terra.Cell(agent.X(), agent.Y()).Energy())
	return 1
}

func (c *EatCommand) handleOmnivore(agent contracts.Agent, terra contracts.Terrain) int {
	whereAddress := agent.Address() + 1
	direction := agent.Command(&whereAddress)
	targetCell := terra.GetNeighbor(agent.X(), agent.Y(), direction)
	if targetCell == nil {
		return 0
	}

	if !targetCell.IsAgent() {
		return 1
	}

	if !slices.Contains([]enum.AgentType{enum.AgentTypeHerbivore, enum.AgentTypeOmnivore, enum.AgentTypePlant, enum.AgentTypeCarnivore}, targetCell.Agent().AgentType()) {
		return 1
	}

	agent.IncreaseEnergy(targetCell.Agent().Energy())
	targetCell.Agent().Kill(terra)
	terra.Cell(agent.X(), agent.Y()).RemoveAgent()
	targetCell.SetAgent(agent)

	return 1
}
