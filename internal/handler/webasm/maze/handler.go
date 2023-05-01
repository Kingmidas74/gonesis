//go:build js && wasm

package maze

import (
	"syscall/js"

	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze"
)

func (h *Handler[G]) Generate(args []js.Value) (*maze.Maze, error) {
	width := args[0].Int()
	height := args[1].Int()

	result, err := h.mazeService.Generate(width, height)
	if err != nil {
		return nil, err
	}

	return result, nil

}
