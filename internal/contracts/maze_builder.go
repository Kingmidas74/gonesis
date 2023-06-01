package contracts

type MazeBuilder interface {
	SetWidth(width int) MazeBuilder
	SetHeight(height int) MazeBuilder
	FirstFilled(flag bool) MazeBuilder
	SetRequiredEmptyCells(requiredEmptyCells int) MazeBuilder
	Build() (Maze, error)
}
