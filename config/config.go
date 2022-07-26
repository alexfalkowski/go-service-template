package config

import (
	"github.com/alexfalkowski/go-service-template/health"
	"github.com/alexfalkowski/go-service/config"
)

// Config for the service.
type Config struct {
	Health        health.Config `yaml:"health"`
	config.Config `yaml:",inline"`
}
