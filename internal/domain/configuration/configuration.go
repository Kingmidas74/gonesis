package configuration

import (
	"encoding/json"
)

type AgentConfiguration struct {
	MaxEnergy            int `json:"MaxEnergy"`
	InitialCount         int `json:"InitialCount"`
	MaxDailyCommandCount int `json:"MaxDailyCommandCount"`
	InitialEnergy        int `json:"InitialEnergy"`
	BrainVolume          int `json:"BrainVolume"`
	MaxSteps             int `json:"MaxSteps"`
}

type Configuration struct {
	PlantConfiguration      AgentConfiguration `json:"PlantConfiguration"`
	HerbivoreConfiguration  AgentConfiguration `json:"HerbivoreConfiguration"`
	CarnivoreConfiguration  AgentConfiguration `json:"CarnivoreConfiguration"`
	DecomposerConfiguration AgentConfiguration `json:"DecomposerConfiguration"`
	OmnivoreConfiguration   AgentConfiguration `json:"OmnivoreConfiguration"`
}

var (
	instance *Configuration
)

func Instance() *Configuration {
	instance = &Configuration{
		PlantConfiguration: AgentConfiguration{
			MaxEnergy:            100,
			InitialCount:         0,
			MaxDailyCommandCount: 1,
			InitialEnergy:        75,
			BrainVolume:          20,
			MaxSteps:             1,
		},
		HerbivoreConfiguration: AgentConfiguration{
			MaxEnergy:            100,
			InitialCount:         0,
			MaxDailyCommandCount: 1,
			InitialEnergy:        75,
			BrainVolume:          20,
			MaxSteps:             1,
		},
		CarnivoreConfiguration: AgentConfiguration{
			MaxEnergy:            100,
			InitialCount:         0,
			MaxDailyCommandCount: 1,
			InitialEnergy:        75,
			BrainVolume:          20,
			MaxSteps:             1,
		},
		DecomposerConfiguration: AgentConfiguration{
			MaxEnergy:            100,
			InitialCount:         0,
			MaxDailyCommandCount: 1,
			InitialEnergy:        75,
			BrainVolume:          20,
			MaxSteps:             1,
		},
		OmnivoreConfiguration: AgentConfiguration{
			MaxEnergy:            100,
			InitialCount:         0,
			MaxDailyCommandCount: 1,
			InitialEnergy:        75,
			BrainVolume:          20,
			MaxSteps:             1,
		},
	}
	return instance
}

func (c *Configuration) FromJson(jsonString string) error {
	return json.Unmarshal([]byte(jsonString), c)
}
