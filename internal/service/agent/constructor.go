package agent

type RandomIntGenerator interface {
	Between(min, max int) int
	Generate(n int) int
	GenerateRandomIntSequence(length int, availableValues []int) []int
}

type RandomFloatGenerator interface {
	Generate() float64
}

type Params struct {
	RandomIntGenerator   RandomIntGenerator
	RandomFloatGenerator RandomFloatGenerator
}

type srv struct {
	randomIntGenerator   RandomIntGenerator
	randomFloatGenerator RandomFloatGenerator
}

func New(params Params) Service {
	return &srv{
		randomIntGenerator:   params.RandomIntGenerator,
		randomFloatGenerator: params.RandomFloatGenerator,
	}
}
