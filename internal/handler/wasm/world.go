//go:build js && wasm

package wasm

import (
	"math/rand"
	"strconv"
	"syscall/js"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/commands"
)

func (h *Handler) InitWorld() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		configJson := args[0].String()

		config, err := h.deserializeConfiguration(configJson)
		if err != nil {
			return h.serializeResponse(1, err.Error())
		}

		seed, err := strconv.ParseInt(config.WorldConfiguration.Seed, 10, 64)
		if err != nil {
			return h.serializeResponse(1, err.Error())
		}

		rand.Seed(seed)

		world, err := h.worldService.Init(config, []contracts.Command{
			commands.NewPhotosynthesisCommand(),
			commands.NewEatCommand(),
			commands.NewMoveCommand(),
			commands.NewPhotosynthesisCommand(),
			commands.NewEatCommand(),
			commands.NewMoveCommand(),
			commands.NewPhotosynthesisCommand(),
			commands.NewEatCommand(),
			commands.NewMoveCommand(),
			commands.NewPhotosynthesisCommand(),
			commands.NewEatCommand(),
			commands.NewMoveCommand(),
			commands.NewPhotosynthesisCommand(),
			commands.NewEatCommand(),
			commands.NewMoveCommand(),
			commands.NewPhotosynthesisCommand(),
			commands.NewEatCommand(),
			commands.NewMoveCommand(),
			commands.NewPhotosynthesisCommand(),
			commands.NewEatCommand(),
			commands.NewMoveCommand(),
			commands.NewPhotosynthesisCommand(),
			commands.NewEatCommand(),
			commands.NewMoveCommand(),
			commands.NewPhotosynthesisCommand(),
			commands.NewEatCommand(),
			commands.NewMoveCommand(),
			commands.NewPhotosynthesisCommand(),
			commands.NewEatCommand(),
			commands.NewMoveCommand(),
			commands.NewCallSubroutineCommand(),
			commands.NewEndSubroutineCommand(),
		})
		if err != nil {
			return h.serializeResponse(1, err.Error())
		}

		return h.serializeWorld(world)
	})
}

func (h *Handler) UpdateWorld() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		configJson := args[0].String()

		config, err := h.deserializeConfiguration(configJson)
		if err != nil {
			return h.serializeResponse(1, err.Error())
		}

		world, err := h.worldService.Update(config)
		if err != nil {
			return h.serializeResponse(1, err.Error())
		}

		return h.serializeWorld(world)
	})
}
