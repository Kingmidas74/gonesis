package maze

import "github.com/kingmidas74/gonesis-engine/internal/contracts"

type Maze struct {
	width   int
	height  int
	content []contracts.Cell
}

func newMaze(width, height int, content []contracts.Cell) contracts.Maze {
	return &Maze{
		width:   width,
		height:  height,
		content: content,
	}
}

func (m Maze) Width() int {
	return m.width
}

func (m Maze) Height() int {
	return m.height
}

func (m Maze) Content() []contracts.Cell {
	return m.content
}
