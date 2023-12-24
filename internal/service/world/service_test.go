package world

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/configuration"
	"log"
	"testing"
)

func TestUpdate(t *testing.T) {
	t.Run("should update the world", func(t *testing.T) {
		config := configuration.NewConfiguration()
		config.WorldConfiguration.Ratio.Width = 11
		config.WorldConfiguration.Ratio.Height = 11
		config.PlantConfiguration.InitialCount = 10
		config.HerbivoreConfiguration.InitialCount = 10
		config.CarnivoreConfiguration.InitialCount = 10
		config.OmnivoreConfiguration.InitialCount = 10
		config.PlantConfiguration.MaxDailyCommandCount = 1
		config.HerbivoreConfiguration.MaxDailyCommandCount = 1
		config.CarnivoreConfiguration.MaxDailyCommandCount = 1
		config.OmnivoreConfiguration.MaxDailyCommandCount = 1

		for i := 0; i < 10000; i++ {
			log.Println("--------------------")
			worldService := New()

			worldService.Init(config)
			for j := 0; j < 1000; j++ {
				worldService.Update(config)
				worldService.Update(config)
			}
		}

	})
}
