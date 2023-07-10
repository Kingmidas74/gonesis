//go:build js && wasm

package wasm

import (
	"math/rand"
	"strconv"
	"syscall/js"

	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
)

func (h *Handler) InitWorld() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		configJson := args[0].String()

		config := configuration.NewConfiguration()
		err := config.FromJson(configJson)
		if err != nil {
			return h.serializeResponse(1, err.Error())
		}

		seed, err := strconv.ParseInt(config.WorldConfiguration.Seed, 10, 64)
		if err != nil {
			return h.serializeResponse(1, err.Error())
		}

		rand.Seed(seed)

		h.gameService.UpdateConfiguration(config)

		world, err := h.gameService.InitWorld()
		if err != nil {
			return h.serializeResponse(1, err.Error())
		}

		return h.serializeWorld(world)
	})
}

func (h *Handler) UpdateWorld() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		world, err := h.gameService.Next()
		if err != nil {
			return h.serializeResponse(1, err.Error())
		}

		return h.serializeWorld(world)
	})
}
