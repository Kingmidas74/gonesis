package call_subroutine

type Command struct {
	isInterrupt bool
}

func New(isInterrupt bool) *Command {
	return &Command{
		isInterrupt: isInterrupt,
	}
}