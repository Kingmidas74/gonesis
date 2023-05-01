package maze

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze"
	"go.uber.org/zap"
	"reflect"
)

func (s Service[K]) Generate(width, height int) (*maze.Maze, error) {
	l := zap.L().Named("maze.generate").
		With(zap.Int("width", width)).
		With(zap.Int("height", height)).
		With(zap.String("generator", reflect.TypeOf(new(K)).String()))

	m, err := maze.NewMazeBuilder[K]().SetHeight(height).SetWidth(width).FirstFilled(true).Build()
	if err != nil {
		l.Error("error", zap.Error(err))
		return nil, err
	}

	return m, nil
}
