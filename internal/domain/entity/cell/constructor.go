package cell

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/contract"
)

type Cell struct {
	x int
	y int

	agent contract.Agent

	northWall bool
	southWall bool
	westWall  bool
	eastWall  bool
}

type WallInfo struct {
	North bool
	South bool
	West  bool
	East  bool
}

func New(x, y int, wallInfo WallInfo) *Cell {
	return &Cell{
		x: x,
		y: y,

		northWall: wallInfo.North,
		southWall: wallInfo.South,
		westWall:  wallInfo.West,
		eastWall:  wallInfo.East,
	}
}
