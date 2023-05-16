package game

import (
	"math/rand"
	"time"

	"golang.org/x/exp/slices"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/commands"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/agent"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze/generator"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/terrain"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/terrain/topology"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/world"
)

func (s *Service) InitWorld(width, height int, agentsCount int) (contracts.World, error) {
	rand.Seed(time.Now().UnixNano())

	mazeBuilder := maze.NewMazeBuilder[generator.SidewinderGenerator]()
	mazeBuilder.SetWidth(width)
	mazeBuilder.SetHeight(height)
	m, err := mazeBuilder.Build()
	if err != nil {
		return nil, err
	}
	terra := terrain.NewTerrain[topology.NeumannTopology](m)
	emptyCells := terra.EmptyCells()
	if len(emptyCells) < agentsCount {
		return nil, ErrNotEnoughEmptyCells
	}
	agents := make([]contracts.Agent, agentsCount)
	availableCommands := []int{0, 1, 2, 3, 4}
	actualCommands := make([]contracts.Command, len(availableCommands))
	for i, _ := range availableCommands {
		actualCommands[i] = commands.NewMoveCommand()
	}

	for i := 0; i < agentsCount; i++ {
		agents[i] = agent.New(rand.Intn(1), s.generateCommandsSequence(availableCommands, 20))
	}

	pickedCellIndexes := make([]int, 0)
	for i := 0; i < agentsCount; i++ {
		targetIndex := rand.Intn(len(emptyCells))
		if slices.Contains(pickedCellIndexes, targetIndex) {
			i--
			continue
		}
		pickedCellIndexes = append(pickedCellIndexes, targetIndex)
		emptyCell := emptyCells[targetIndex]
		agents[i].SetX(emptyCell.X())
		agents[i].SetY(emptyCell.Y())
	}

	s.world = world.New(terra, agents, actualCommands)
	return s.world, nil
}

func (s *Service) generateCommandsSequence(availableCommands []int, sequanceLength int) []int {
	result := make([]int, sequanceLength)
	for i := 0; i < sequanceLength; i++ {
		// Pick a random index into availableCommands
		index := rand.Intn(len(availableCommands))
		// Add the chosen command to the result
		result[i] = availableCommands[index]
	}
	return result
}
