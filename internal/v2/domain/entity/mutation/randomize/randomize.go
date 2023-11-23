package randomize

type Brain interface {
	Commands() []int
	Address() int
}

type RandomIntGenerator interface {
	Intn(n int) int
}

type Mutation struct {
	randomIntGenerator RandomIntGenerator
}

func New(randomIntGenerator RandomIntGenerator) *Mutation {
	return &Mutation{
		randomIntGenerator: randomIntGenerator,
	}
}

func (b Mutation) Mutate(subject Brain) error {
	commands := subject.Commands()
	if len(commands) <= 0 {
		return nil
	}

	commands[subject.Address()] = b.randomIntGenerator.Intn(len(commands))
	return nil
}
