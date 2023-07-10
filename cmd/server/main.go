package main

import (
	"log"

	"go.uber.org/zap"

	"github.com/kingmidas74/gonesis-engine/internal/app"
	"github.com/kingmidas74/gonesis-engine/internal/config"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	cfg, err := config.New()
	if err != nil {
		zap.L().Error("Failed to load configuration", zap.Error(err))
		return
	}

	a, err := app.New(cfg, logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	err = a.Init()
	if err != nil {
		logger.Fatal(err.Error())
	}

	err = a.Run()
	if err != nil {
		logger.Fatal(err.Error())
	}
}
