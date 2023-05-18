package model

type World struct {
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Cells  []Cell  `json:"cells"`
	Agents []Agent `json:"agents"`
}

type Cell struct {
	CellType int `json:"cellType"`
}

type Agent struct {
	Brain
	X      int `json:"x"`
	Y      int `json:"y"`
	Energy int `json:"energy"`
}

type Brain struct {
	Commands []int `json:"commands"`
}
