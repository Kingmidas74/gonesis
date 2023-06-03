package contracts

import "github.com/kingmidas74/gonesis-engine/internal/domain/enum"

type ReproductionSystem interface {
	ReproductionType() enum.ReproductionSystemType
	Reproduce(parents []Agent) ([]Agent, error)
}
