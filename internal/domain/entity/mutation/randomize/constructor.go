package randomize

type RandomIntGenerator interface {
	Generate(n int) int
}

type Mutation struct {
	randomIntGenerator RandomIntGenerator
}

func New(randomIntGenerator RandomIntGenerator) *Mutation {
	return &Mutation{
		randomIntGenerator: randomIntGenerator,
	}
}
