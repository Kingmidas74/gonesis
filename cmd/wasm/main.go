//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	domain_maze "github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze"
	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze/generator"
	"github.com/kingmidas74/gonesis-engine/internal/service/maze"
)

func main() {
	js.Global().Set("generateSideWinderMaze", generateMaze[generator.SidewinderGenerator]())
	js.Global().Set("generateAldousBroderMaze", generateMaze[generator.AldousBroderGenerator]())
	js.Global().Set("generateBinaryMaze", generateMaze[generator.BinaryGenerator]())
	<-make(chan bool)
}

func generateMaze[G domain_maze.Generator]() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("here 0")
		mazeService := maze.NewMazeService[G]()
		fmt.Println("here 1")

		result, err := mazeService.Generate(5, 5)
		fmt.Println("here 2")
		if err != nil {
			fmt.Println("here 3")
			return err.Error()
		}
		fmt.Println("here 4")
		fmt.Println(result)
		fmt.Println("here 5")

		res, _ := json.Marshal(result)
		fmt.Println("here 6")
		fmt.Println(string(res))
		panic("ASD")
		return string(res)
		/*
			resultValue := reflect.ValueOf(*result)
			if resultValue.Kind() == reflect.Ptr {
				resultValue = resultValue.Elem()
			}

			resultMap := make(map[string]interface{})
			resultType := resultValue.Type()

			for i := 0; i < resultValue.NumField(); i++ {
				fieldName := resultType.Field(i).Name
				fieldValue := resultValue.Field(i).Interface()
				resultMap[fieldName] = fieldValue
			}

			return js.ValueOf(resultMap)

		*/
	})
}
