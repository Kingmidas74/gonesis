//go:build js && wasm

package game

import (
	"encoding/json"
	"github.com/kingmidas74/gonesis-engine/internal/mapper"
	"syscall/js"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/handler/webasm/model"
)

func (h *Handler) InitWorldAndRun(args []js.Value, callback func(int, string)) (*model.World, error) {
	width := args[0].Int()
	height := args[1].Int()
	agentsCount := args[2].Int()

	world, err := h.gameService.InitWorld(width, height, agentsCount)
	if err != nil {
		return nil, err
	}

	result := mapper.NewWorld(world)

	world.Action(1, func(w contracts.World, currentDay int) {
		if r, e := json.Marshal(mapper.NewWorld(w)); e == nil {
			callback(currentDay, string(r))
		}
	})
	return &result, nil
}

func (h *Handler) InitWorld(args []js.Value, callback func(int, string)) (*model.World, error) {
	width := args[0].Int()
	height := args[1].Int()
	agentsCount := args[2].Int()

	world, err := h.gameService.InitWorld(width, height, agentsCount)
	if err != nil {
		return nil, err
	}

	result := mapper.NewWorld(world)

	world.Action(1, func(w contracts.World, currentDay int) {
		if r, e := json.Marshal(mapper.NewWorld(w)); e == nil {
			callback(currentDay, string(r))
		}
	})
	return &result, nil
}
