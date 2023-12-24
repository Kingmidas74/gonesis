package moore

import "github.com/kingmidas74/gonesis-engine/pkg/topology/2d/moore"

type Topology struct {
	topology moore.Topology
}

func New() *Topology {
	return &Topology{
		topology: moore.New(),
	}
}
