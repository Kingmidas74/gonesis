package configuration

import (
	"encoding/json"
	"github.com/kingmidas74/gonesis-engine/internal/domain/enum"
)

const (
	defaultMaxEnergy     = 100
	defaultBrainSize     = 20
	defaultEnergy        = 75
	defaultCount         = 0
	defaultDailyCommands = 1
)

type WorldConfiguration struct {
	MazeType enum.MazeType     `json:"MazeType"`
	Topology enum.TopologyType `json:"Topology"`
}

type AgentConfiguration struct {
	MaxEnergy            int `json:"MaxEnergy"`
	InitialCount         int `json:"InitialCount"`
	MaxDailyCommandCount int `json:"MaxDailyCommandCount"`
	InitialEnergy        int `json:"InitialEnergy"`
	BrainVolume          int `json:"BrainVolume"`
}

type Configuration struct {
	WorldConfiguration      WorldConfiguration `json:"WorldConfiguration"`
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
		WorldConfiguration: WorldConfiguration{
			MazeType: enum.MazeTypeEmpty,
			Topology: enum.TopologyTypeNeumann,
		},
		PlantConfiguration: AgentConfiguration{
			MaxEnergy:            defaultMaxEnergy,
			InitialCount:         defaultCount,
			MaxDailyCommandCount: defaultDailyCommands,
			InitialEnergy:        defaultEnergy,
			BrainVolume:          defaultBrainSize,
		},
		HerbivoreConfiguration: AgentConfiguration{
			MaxEnergy:            defaultMaxEnergy,
			InitialCount:         defaultCount,
			MaxDailyCommandCount: defaultDailyCommands,
			InitialEnergy:        defaultEnergy,
			BrainVolume:          defaultBrainSize,
		},
		CarnivoreConfiguration: AgentConfiguration{
			MaxEnergy:            defaultMaxEnergy,
			InitialCount:         defaultCount,
			MaxDailyCommandCount: defaultDailyCommands,
			InitialEnergy:        defaultEnergy,
			BrainVolume:          defaultBrainSize,
		},
		DecomposerConfiguration: AgentConfiguration{
			MaxEnergy:            defaultMaxEnergy,
			InitialCount:         defaultCount,
			MaxDailyCommandCount: defaultDailyCommands,
			InitialEnergy:        defaultEnergy,
			BrainVolume:          defaultBrainSize,
		},
		OmnivoreConfiguration: AgentConfiguration{
			MaxEnergy:            defaultMaxEnergy,
			InitialCount:         defaultCount,
			MaxDailyCommandCount: defaultDailyCommands,
			InitialEnergy:        defaultEnergy,
			BrainVolume:          defaultBrainSize,
		},
	}
	return instance
}

func (c *Configuration) FromJson(jsonString string) error {
	return json.Unmarshal([]byte(jsonString), c)
}
