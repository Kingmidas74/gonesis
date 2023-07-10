package config

import (
	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigyaml"
)

type Config struct {
	Host Host `yaml:"host"`
}

type Host struct {
	Port         string `yaml:"port" env:"PORT" default:"9091"`
	StaticFolder string `yaml:"static_folder" env:"STATIC_FOLDER" default:"./web"`
}

func New() (Config, error) {
	var cfg Config

	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		AllowUnknownEnvs: false,
		SkipFlags:        true,
		EnvPrefix:        "GONESIS",
		Files:            []string{"./config/config.yaml"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".yaml": aconfigyaml.New(),
		},
	})

	if err := loader.Load(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
