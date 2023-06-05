package model

type World struct {
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Cells  []Cell  `json:"cells"`
	Agents []Agent `json:"agents"`
}

type Cell struct {
	CellType string `json:"cellType"`
	Energy   int    `json:"energy"`
}

type Agent struct {
	Brain
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Energy    int    `json:"energy"`
	AgentType string `json:"agentType"`
}

type Brain struct {
	Commands []int `json:"commands"`
}