package model

type World struct {
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Cells  []Cell  `json:"cells"`
	Agents []Agent `json:"agents"`
}

type Cell struct {
	CellType uint8 `json:"cellType"`
}

type Agent struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Energy int `json:"energy"`
}
