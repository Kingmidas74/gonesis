//go:build js && wasm

package maze

import (
	"syscall/js"

	"github.com/kingmidas74/gonesis-engine/internal/handler/webasm/model"
)

func (h *Handler[G]) Generate(args []js.Value) (*model.Maze, error) {
	width := args[0].Int()
	height := args[1].Int()

	maze, err := h.mazeService.Generate(width, height)
	if err != nil {
		return nil, err
	}

	return &model.Maze{
		Width:   maze.Width(),
		Height:  maze.Height(),
		Content: maze.Content(),
	}, nil
}
