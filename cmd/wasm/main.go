//go:build js && wasm

package main

import (
	"encoding/json"
	"github.com/kingmidas74/gonesis-engine/internal/handler/webasm/game"
	"math/rand"
	"syscall/js"
	"time"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze/generator"
	"github.com/kingmidas74/gonesis-engine/internal/handler/webasm/maze"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	js.Global().Set("generateSideWinderMaze", generateMaze[generator.SidewinderGenerator]())
	js.Global().Set("generateAldousBroderMaze", generateMaze[generator.AldousBroderGenerator]())
	js.Global().Set("generateBinaryMaze", generateMaze[generator.BinaryGenerator]())
	js.Global().Set("generateGridMaze", generateMaze[generator.GridGenerator]())
	js.Global().Set("generateBorder", generateMaze[generator.BorderGenerator]())
	js.Global().Set("updateState", updateState())
	js.Global().Set("runGame", runGame())
	<-make(chan bool)
}

func generateMaze[G contracts.MazeGenerator]() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		h := maze.New[G]()
		res, err := h.Generate(args)
		if err != nil {
			return err.Error()
		}

		js.Global().Call("fromGo", "etes")
		if r, e := json.Marshal(res); e != nil {
			return e.Error()
		} else {
			return string(r)
		}
	})
}

func updateState() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return js.ValueOf(make([]bool, 0))
	})
}

func dayCallback(day int, jsonData string) {
	js.Global().Call("fromGo", day, jsonData)
}

func runGame() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		h := game.New()
		//res, err := h.InitWorld(args)
		res, err := h.InitWorldAndRun(args, dayCallback)
		if err != nil {
			return err.Error()
		}
		if r, err := json.Marshal(res); err == nil {
			return string(r)
		}
		return err.Error()
	})
}
