package cmd

import (
	"github.com/alexfalkowski/go-service-template/internal/config"
	"github.com/alexfalkowski/go-service-template/internal/health"
	"github.com/alexfalkowski/go-service/v2/module"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	module.Server,
	config.Module,
	health.Module,
)
