package contracts

type Energy interface {
	Energy() int
	IncreaseEnergy(delta int)
	DecreaseEnergy(delta int)
}
