//go:build js && wasm

package main

import (
	"encoding/json"
	"math/rand"
	"syscall/js"
	"time"

	domain_maze "github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze"
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
	<-make(chan bool)
}

func generateMaze[G domain_maze.Generator]() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		h := maze.New[G]()
		res, err := h.Generate(args)
		if err != nil {
			return err.Error()
		}

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
