package world

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/domain/contract"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/cell"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/terrain"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/topology/moore"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/world"
)

func (s *srv) Init(config *configuration.Configuration) (contract.World, error) {
	mazeGenerator, err := s.mazeGeneratorCollection.Get(config.WorldConfiguration.MazeType)
	if err != nil {
		return nil, err
	}
	cells, err := mazeGenerator.Generate(config.WorldConfiguration.Ratio.Width, config.WorldConfiguration.Ratio.Height)
	if err != nil {
		return nil, err
	}

	terrainCells := make([]contract.Cell, len(cells))
	for i, c := range cells {
		terrainCells[i] = cell.New(c.X(), c.Y(), cell.WallInfo{
			North: c.NorthWall(),
			East:  c.EastWall(),
			South: c.SouthWall(),
			West:  c.WestWall(),
		})
	}

	topology := moore.New()
	terra := terrain.New(terrain.MazeInfo{
		Width:  config.WorldConfiguration.Ratio.Width,
		Height: config.WorldConfiguration.Ratio.Height,
		Cells:  terrainCells,
	}, terrain.WithTopology(topology))

	agents, err := s.agentService.Generate(config)
	if err != nil {
		return nil, err
	}

	w := world.New(world.WithTerrain(terra))

	err = w.PlaceAgents(agents)
	if err != nil {
		return nil, err
	}

	s.world = w
	return w, nil
}

func (s *srv) Update(config *configuration.Configuration) (contract.World, error) {
	err := s.world.Next()
	if err != nil {
		return nil, err
	}
	return s.world, nil
}
