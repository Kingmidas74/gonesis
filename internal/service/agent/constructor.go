package agent

type srv struct {
}

func New() Service {
	return &srv{}
}
