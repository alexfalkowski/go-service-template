package config

import (
	"github.com/alexfalkowski/go-service-template/health"
	"github.com/alexfalkowski/go-service/config"
)

// Config for the service.
type Config struct {
	Health         *health.Config `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}
