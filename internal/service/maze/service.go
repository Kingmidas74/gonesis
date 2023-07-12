package maze

import (
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze/generator"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
	"github.com/kingmidas74/gonesis-engine/internal/domain/errors"
)

func (s *srv) Generate(mazeType enum.MazeType, width, height, requiredEmptyCells int) (contracts.Maze, error) {
	mazeBuilder, err := s.createBuilderByType(mazeType)
	if err != nil {
		return nil, err
	}

	m, err := mazeBuilder.SetWidth(width).
		SetHeight(height).
		SetRequiredEmptyCells(requiredEmptyCells).
		Build()

	return m, err
}

func (s *srv) createBuilderByType(mazeType enum.MazeType) (contracts.MazeBuilder, error) {
	switch mazeType {
	case enum.MazeTypeBorder:
		return maze.NewMazeBuilder[generator.BorderGenerator](), nil
	case enum.MazeTypeBinary:
		return maze.NewMazeBuilder[generator.BinaryGenerator](), nil
	case enum.MazeTypeGrid:
		return maze.NewMazeBuilder[generator.GridGenerator](), nil
	case enum.MazeTypeAldousBroder:
		return maze.NewMazeBuilder[generator.AldousBroderGenerator](), nil
	case enum.MazeTypeSideWinder:
		return maze.NewMazeBuilder[generator.SidewinderGenerator](), nil
	case enum.MazeTypeEmpty:
		return maze.NewMazeBuilder[generator.EmptyGenerator](), nil
	default:
		return nil, errors.ErrMazeTypeNotSupported
	}
}
