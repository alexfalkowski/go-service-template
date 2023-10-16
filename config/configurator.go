package config

import (
	"github.com/alexfalkowski/go-service-template/health"
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/config"
)

// NewConfigurator for config.
func NewConfigurator(i *cmd.InputConfig) (config.Configurator, error) {
	c := &Config{}

	if err := i.Unmarshal(c); err != nil {
		return nil, err
	}

	return c, nil
}

func healthConfig(cfg config.Configurator) *health.Config {
	return &cfg.(*Config).Health
}
