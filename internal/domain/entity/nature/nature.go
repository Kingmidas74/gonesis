package nature

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/contract"
)

func (n *Nature) BrainVolume() int {
	return n.brainVolume
}

func (n *Nature) ReproductionEnergyCost() int {
	return n.reproductionEnergyCost
}

func (n *Nature) ReproductionChance() float64 {
	return n.reproductionChance
}

func (n *Nature) MutationChance() float64 {
	return n.mutationChance
}

func (n *Nature) MaxDailyCommandCount() int {
	return n.maxDailyCommandCount
}

func (n *Nature) FindCommand(identifier int) contract.Command {
	if identifier < 0 || identifier >= len(n.commands) {
		return nil
	}
	return n.commands[identifier]
}

func (n *Nature) InitialEnergy() int {
	return n.initialEnergy
}

func (n *Nature) Type() contract.AgentType {
	return n.agentType
}

func (a *Nature) AvailableFood() map[contract.AgentType]int {
	return map[contract.AgentType]int{
		contract.AgentTypeHerbivore: 8,
		contract.AgentTypeOmnivore:  8,
	}
}

func (a *Nature) Reproduction() contract.Reproduction {
	return a.reproductionSystem
}

func (a *Nature) InitialCount() int {
	return a.initialCount
}
