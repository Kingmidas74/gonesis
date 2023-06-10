package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"golang.org/x/net/http2"
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

	fs := http.FileServer(http.Dir(assetsDir))

	http.Handle("/", noCache(fs))

	srv := &http.Server{
		Addr:    ":9091",
		Handler: http.FileServer(http.Dir(assetsDir)),
	}

	if err := http2.ConfigureServer(srv, &http2.Server{}); err != nil {
		fmt.Println("Failed to configure server:", err)
		return
	}

	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
}

func noCache(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		h.ServeHTTP(w, r)
	}
}
