package contracts

type MazeGenerator interface {
	Generate(width, height int) (maze []bool, err error)
}
