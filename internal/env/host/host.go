package host

import (
	"github.com/kingmidas74/gonesis-engine/internal/config"
)

type Host struct {
	Port         string
	StaticFolder string
}

func New(conf config.Host) (*Host, error) {
	return &Host{
		Port:         conf.Port,
		StaticFolder: conf.StaticFolder,
	}, nil
}
