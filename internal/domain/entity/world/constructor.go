package world

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/contract"
)

type World struct {
	terrain contract.Terrain

	currentDay int
}

func New(options ...func(*World)) *World {
	w := &World{
		currentDay: 0,
	}
	for _, o := range options {
		o(w)
	}
	return w
}

func WithTerrain(terrain contract.Terrain) func(*World) {
	return func(w *World) {
		w.terrain = terrain
	}
}
