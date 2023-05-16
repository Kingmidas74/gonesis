//go:build js && wasm

package main

import (
	"encoding/json"
	"github.com/kingmidas74/gonesis-engine/internal/mapper"
	"syscall/js"

	"github.com/kingmidas74/gonesis-engine/internal/contracts"
	"github.com/kingmidas74/gonesis-engine/internal/service/game"
)

func main() {
	js.Global().Set("initWorld", initWorld())
	js.Global().Set("step", step())

	/*
		js.Global().Set("generateSideWinderMaze", generateMaze[generator.SidewinderGenerator]())
		js.Global().Set("generateAldousBroderMaze", generateMaze[generator.AldousBroderGenerator]())
		js.Global().Set("generateBinaryMaze", generateMaze[generator.BinaryGenerator]())
		js.Global().Set("generateGridMaze", generateMaze[generator.GridGenerator]())
		js.Global().Set("generateBorder", generateMaze[generator.BorderGenerator]())
		js.Global().Set("updateState", updateState())
		js.Global().Set("runGame", runGame())
	*/

	<-make(chan bool)
}

var GameWorld contracts.World

func initWorld() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		width := args[0].Int()
		height := args[1].Int()
		agentsCount := args[2].Int()

		gameService := game.New()
		world, err := gameService.InitWorld(width, height, agentsCount)
		if err != nil {
			return err.Error()
		}

		GameWorld = world

		return serializeWorld(world)
	})
}

func step() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		err := GameWorld.Next()
		if err != nil {
			return err.Error()
		}

		return serializeWorld(GameWorld)
	})
}

func serializeWorld(w contracts.World) string {
	res := mapper.NewWorld(w)

	if r, e := json.Marshal(res); e != nil {
		return e.Error()
	} else {
		return string(r)
	}
}

//js.Global().Call("fromGo", "etes")

/*
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
*/
