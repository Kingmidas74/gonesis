package maze

type Cell interface {
}

type Maze struct {
	width   int
	height  int
	content []Cell
}

func New(width, height int, content []Cell) *Maze {
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

func (m Maze) Content() []Cell {
	return m.content
}
