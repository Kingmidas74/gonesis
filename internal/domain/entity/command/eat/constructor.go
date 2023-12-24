package eat

type Command struct {
	isInterrupt bool
}

func New(isInterrupt bool) *Command {
	return &Command{
		isInterrupt: isInterrupt,
	}
}
