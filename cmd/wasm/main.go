//go:build js && wasm

package main

import (
	"github.com/kingmidas74/gonesis-engine/internal/service/agent"
	"github.com/kingmidas74/gonesis-engine/internal/service/world"
	"syscall/js"

	"github.com/kingmidas74/gonesis-engine/internal/handler/wasm"
	randomize_float "github.com/kingmidas74/gonesis-engine/pkg/randomize/float"
	randomize_int "github.com/kingmidas74/gonesis-engine/pkg/randomize/int"
)

var WasmHandler *wasm.Handler

func main() {

	randomIntGenerator := randomize_int.New()
	randomFloatGenerator := randomize_float.New()

	agentService := agent.New(agent.Params{
		RandomIntGenerator:   randomIntGenerator,
		RandomFloatGenerator: randomFloatGenerator,
	})

	worldService := world.New(world.Params{
		AgentService: agentService,
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
