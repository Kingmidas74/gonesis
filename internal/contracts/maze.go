package contracts

type Maze interface {
	Width() int
	Height() int
	Content() []Cell
}
