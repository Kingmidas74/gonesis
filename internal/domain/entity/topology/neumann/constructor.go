package neumann

import "github.com/kingmidas74/gonesis-engine/pkg/topology/2d/neumann"

type Topology struct {
	topology neumann.Topology
}

func New() *Topology {
	return &Topology{
		topology: neumann.New(),
	}
}
