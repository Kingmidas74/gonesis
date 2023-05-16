package maze

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/errors"
	"math/rand"
)

type Builder[G contracts.MazeGenerator] struct {
	width              int
	height             int
	firstFilled        bool
	requiredEmptyCells int
}

func NewMazeBuilder[G contracts.MazeGenerator]() contracts.MazeBuilder[G] {
	return &Builder[G]{}
}

func (b *Builder[G]) SetWidth(width int) contracts.MazeBuilder[G] {
	b.width = width

	return b
}

func (b *Builder[G]) SetHeight(height int) contracts.MazeBuilder[G] {
	b.height = height

	return b
}

func (b *Builder[G]) FirstFilled(flag bool) contracts.MazeBuilder[G] {
	b.firstFilled = flag

	return b
}

func (b *Builder[G]) SetRequiredEmptyCells(requiredEmptyCells int) contracts.MazeBuilder[G] {
	b.requiredEmptyCells = requiredEmptyCells

	return b
}

func (b *Builder[G]) Build() (contracts.Maze, error) {
	g := *new(G)

	maze, err := g.Generate(b.width, b.height)
	if err != nil {
		return nil, err
	}

	if b.width*b.height < b.requiredEmptyCells {
		return nil, errors.ErrFreeRequirementIncorrect
	}

	filledCells := make([]int, 0)
	for i, c := range maze {
		if !c {
			filledCells = append(filledCells, i)
		}
	}

	for emptyCellCount := b.width*b.height - len(filledCells); emptyCellCount < b.requiredEmptyCells; emptyCellCount++ {
		randIndex := rand.Intn(len(filledCells))
		cellIndex := filledCells[rand.Intn(len(filledCells))]
		filledCells = append(filledCells[:randIndex], filledCells[randIndex+1:]...)
		maze[cellIndex] = true
	}

	return newMaze(b.width, b.height, maze), nil
}
