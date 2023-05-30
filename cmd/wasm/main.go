//go:build js && wasm

package main

import (
	"encoding/json"
	"math/rand"
	"syscall/js"
	"time"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"github.com/kingmidas74/gonesis-engine/internal/handler/webasm/model"
	"github.com/kingmidas74/gonesis-engine/internal/mapper"
	"github.com/kingmidas74/gonesis-engine/internal/service/game"
)

func main() {
	js.Global().Set("initWorld", initWorld())
	js.Global().Set("step", step())

	<-make(chan bool)
}

var GameWorld contracts.World

func initWorld() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		width := args[0].Int()
		height := args[1].Int()
		configJson := args[2].String()
		cfg := configuration.Instance()

		err := cfg.FromJson(configJson)
		if err != nil {
			return serializeError(err)
		}

		rand.Seed(time.Now().UnixNano())

		gameService := game.New(cfg)
		world, err := gameService.InitWorld(width, height)
		if err != nil {
			return serializeError(err)
		}

		GameWorld = world

		return serializeWorld(world)
	})
}

func step() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		err := GameWorld.Next(configuration.Instance())
		if err != nil {
			return serializeError(err)
		}

		return serializeWorld(GameWorld)
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
