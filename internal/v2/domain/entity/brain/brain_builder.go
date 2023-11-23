package brain

import "errors"

type CommandSequenceGenerator interface {
	Generate(volume int) ([]int, error)
}

type Builder struct {
	brain *Brain
}

func NewBuilder() *Builder {
	return &Builder{
		brain: &Brain{
			address: 0,
		},
	}
}

func (b *Builder) WithCommands(commands []int) *Builder {
	b.brain.commands = commands
	return b
}

func (b *Builder) WithVolume(volume int, generator CommandSequenceGenerator) (*Builder, error) {
	if volume <= 0 {
		return b, ErrBrainVolumeIncorrect
	}
	var err error
	b.brain.commands, err = generator.Generate(volume)
	if err != nil {
		return b, err
	}
	return b, nil
}

func (b *Builder) Build() *Brain {
	return b.brain
}

var ErrBrainVolumeIncorrect = errors.New("brain's volume is incorrect ")
