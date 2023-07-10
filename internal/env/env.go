package env

import (
	"fmt"

	"github.com/kingmidas74/gonesis-engine/internal/config"
	"github.com/kingmidas74/gonesis-engine/internal/env/host"
)

type Env struct {
	Host *host.Host
}

func New(cfg config.Config) (*Env, error) {
	hostEnv, err := host.New(cfg.Host)
	if err != nil {
		return nil, fmt.Errorf("failed to create a host env: %s", err)
	}

	return &Env{
		Host: hostEnv,
	}, nil
}
