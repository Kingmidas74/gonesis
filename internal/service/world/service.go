package world

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	contract "github.com/kingmidas74/gonesis-engine/internal/domain/contract"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/cell"
	"github.com/kingmidas74/gonesis-engine/pkg/graph_maze/empty"
	"github.com/kingmidas74/gonesis-engine/pkg/graph_topology/neumann"
)

func (s *srv) Init(config *configuration.Configuration) (contract.World, error) {
	ng := neumann.New[cell.Cell]()
	g, err := ng.Generate(config.WorldConfiguration.Ratio.Width, config.WorldConfiguration.Ratio.Height)
	if err != nil {
		return nil, err
	}

	mg := empty.New[cell.Cell, neumann.VertexID]()
	mg.Generate(g, g.FindVertex(neumann.VertexID{X: 0, Y: 0}))

	/*
		cells, err := empty.New().Generate(config.WorldConfiguration.Ratio.Width, config.WorldConfiguration.Ratio.Height)
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

	*/
}

func (s *srv) Update(config *configuration.Configuration) (contract.World, error) {
	err := s.world.Next()
	if err != nil {
		return nil, err
	}
	return s.world, nil
}
