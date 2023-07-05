//go:build js && wasm

package main

import (
	"encoding/json"
	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/mapper"
	"github.com/kingmidas74/gonesis-engine/internal/mapper/model"
	"github.com/kingmidas74/gonesis-engine/internal/service/game"
	"math/rand"
	"strconv"
	"syscall/js"
)

func main() {
	js.Global().Set("initWorld", initWorld())
	js.Global().Set("step", step())

	<-make(chan bool)
}

var GameService *game.Service

func initWorld() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		configJson := args[0].String()

		config := configuration.NewConfiguration()
		err := config.FromJson(configJson)
		if err != nil {
			return serializeError(err)
		}

		seed, _ := strconv.ParseInt(config.WorldConfiguration.Seed, 10, 64)
		rand.Seed(seed)

		GameService = game.New(config)
		world, err := GameService.InitWorld()
		if err != nil {
			return serializeError(err)
		}

		return serializeWorld(world)
	})
}

func step() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		world, err := GameService.Next()
		if err != nil {
			return serializeError(err)
		}

		return serializeWorld(world)
	})
}

func serializeWorld(w contracts.World) string {
	res := mapper.NewWorld(w)

	if r, e := json.Marshal(res); e != nil {
		return serializeError(e)
	} else {
		return serializeResponse(string(r))
	}
}

func serializeError(e error) string {
	r, err := json.Marshal(model.Response{
		Code:    1,
		Message: e.Error(),
	})
	if err != nil {
		return err.Error()
	}
	return string(r)
}

func serializeResponse(message string) string {
	r, err := json.Marshal(model.Response{
		Code:    0,
		Message: message,
	})
	if err != nil {
		return err.Error()
	}
	return string(r)
}
