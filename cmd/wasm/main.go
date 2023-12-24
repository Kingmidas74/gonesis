//go:build js && wasm

package main

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/contract"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze_generator_collection"
	"github.com/kingmidas74/gonesis-engine/internal/service/agent"
	"github.com/kingmidas74/gonesis-engine/internal/service/world"
	"github.com/kingmidas74/gonesis-engine/pkg/maze/generator/aldous_broder"
	"github.com/kingmidas74/gonesis-engine/pkg/maze/generator/binary"
	"github.com/kingmidas74/gonesis-engine/pkg/maze/generator/border"
	"github.com/kingmidas74/gonesis-engine/pkg/maze/generator/empty"
	"github.com/kingmidas74/gonesis-engine/pkg/maze/generator/grid"
	"github.com/kingmidas74/gonesis-engine/pkg/maze/generator/sidewinder"
	"syscall/js"

	"github.com/kingmidas74/gonesis-engine/internal/handler/wasm"
	randomize_float "github.com/kingmidas74/gonesis-engine/pkg/randomize/float"
	randomize_int "github.com/kingmidas74/gonesis-engine/pkg/randomize/int"
)

var WasmHandler *wasm.Handler

func main() {

	randomIntGenerator := randomize_int.New()
	randomFloatGenerator := randomize_float.New()

	mazeGeneratorCollection := maze_generator_collection.New()
	mazeGeneratorCollection.Register(contract.MazeTypeAldousBroder, aldous_broder.New(aldous_broder.Params{RandomIntGenerator: randomIntGenerator}))
	mazeGeneratorCollection.Register(contract.MazeTypeBinary, binary.New(binary.Params{RandomIntGenerator: randomIntGenerator}))
	mazeGeneratorCollection.Register(contract.MazeTypeGrid, grid.New())
	mazeGeneratorCollection.Register(contract.MazeTypeBorder, border.New())
	mazeGeneratorCollection.Register(contract.MazeTypeSideWinder, sidewinder.New(sidewinder.Params{RandomIntGenerator: randomIntGenerator}))
	mazeGeneratorCollection.Register(contract.MazeTypeEmpty, empty.New())

	agentService := agent.New(agent.Params{
		RandomIntGenerator:   randomIntGenerator,
		RandomFloatGenerator: randomFloatGenerator,
	})

	worldService := world.New(world.Params{
		AgentService:            agentService,
		MazeGeneratorCollection: mazeGeneratorCollection,
	})

	WasmHandler, _ = wasm.New(wasm.Params{
		WorldService: worldService,
	})

	js.Global().Set("initWorld", initWorld())
	js.Global().Set("updateWorld", updateWorld())

	<-make(chan bool)
}

func initWorld() js.Func {
	return WasmHandler.InitWorld()
}

func updateWorld() js.Func {
	return WasmHandler.UpdateWorld()
}
