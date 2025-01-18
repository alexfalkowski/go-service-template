package cmd

import (
	"github.com/alexfalkowski/go-service-template/config"
	"github.com/alexfalkowski/go-service-template/health"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/transport"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	module.Module, debug.Module, feature.Module,
	telemetry.Module, transport.Module,
	config.Module, health.Module, Module,
}
