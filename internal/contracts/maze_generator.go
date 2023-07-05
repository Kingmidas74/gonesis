package contracts

type MazeGenerator interface {
	Generate(width, height int) (maze []Cell, err error)
}
