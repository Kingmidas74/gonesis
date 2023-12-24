package model

type World struct {
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Cells      []Cell `json:"cells"`
	CurrentDay int    `json:"currentDay"`
}

type Cell struct {
	Energy    int    `json:"energy"`
	Agent     *Agent `json:"agent"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	NorthWall bool   `json:"northWall"`
	SouthWall bool   `json:"southWall"`
	WestWall  bool   `json:"westWall"`
	EastWall  bool   `json:"eastWall"`
}

type Agent struct {
	Brain
	X          int    `json:"x"`
	Y          int    `json:"y"`
	Energy     int    `json:"energy"`
	AgentType  string `json:"agentType"`
	Generation int    `json:"generation"`
}

type Brain struct {
	Commands []int `json:"commands"`
}
