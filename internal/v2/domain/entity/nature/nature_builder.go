package nature

type Builder struct {
	nature *Nature
}

func NewBuilder() *Builder {
	return &Builder{
		nature: New(),
	}
}

func (b *Builder) WithInitialEnergy(initialEnergy int) *Builder {
	b.nature.initialEnergy = initialEnergy
	return b
}

func (b *Builder) WithBrainVolume(brainVolume int) *Builder {
	b.nature.brainVolume = brainVolume
	return b
}

func (b *Builder) WithReproductionEnergyCost(reproductionEnergyCost int) *Builder {
	b.nature.reproductionEnergyCost = reproductionEnergyCost
	return b
}

func (b *Builder) WithReproductionChance(reproductionChance int) *Builder {
	b.nature.reproductionChance = reproductionChance
	return b
}

func (b *Builder) WithMutationChance(mutationChance int) *Builder {
	b.nature.mutationChance = mutationChance
	return b
}

func (b *Builder) Build() *Nature {
	return b.nature
}
