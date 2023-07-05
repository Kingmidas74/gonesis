package model

type Maze struct {
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Content []bool `json:"content"`
}
