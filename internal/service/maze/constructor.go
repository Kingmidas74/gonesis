package maze

import "github.com/kingmidas74/gonesis-engine/internal/domain/configuration"

type srv struct {
	config *configuration.Configuration
}

func New(config *configuration.Configuration) Service {
	return &srv{
		config: config,
	}
}
