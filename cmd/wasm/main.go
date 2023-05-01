//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	js.Global().Set("example", example())
	<-make(chan bool)
}

func example() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println("example 2")
		return nil
	})
	return jsonFunc
}
