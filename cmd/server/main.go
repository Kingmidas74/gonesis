package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current working directory:", err)
		return
	}

	assetsDir := filepath.Join(cwd, "web")
	fmt.Println("Serving files from:", assetsDir)

	err = http.ListenAndServe(":9091", http.FileServer(http.Dir(assetsDir)))
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
}
