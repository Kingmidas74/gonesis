//go:build js && wasm

package main

import (
	"syscall/js"

	"github.com/kingmidas74/gonesis-engine/internal/handler/wasm"
)

var WasmHandler *wasm.Handler

func main() {

	WasmHandler, _ = wasm.New()

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
