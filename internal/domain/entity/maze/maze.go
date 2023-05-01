package maze

type Maze struct {
	Width   int
	Height  int
	Content []bool
}

func newMaze(width, height int, content []bool) *Maze {
	return &Maze{
		Width:   width,
		Height:  height,
		Content: content,
	}
}

func (m Maze) GetWidth() int {
	return m.Width
}

func (m Maze) GetHeight() int {
	return m.Height
}
