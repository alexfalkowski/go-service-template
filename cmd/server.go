package cmd

import (
	"github.com/alexfalkowski/go-service-template/config"
	"github.com/alexfalkowski/go-service-template/server/health"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/telemetry/metrics"
	"github.com/alexfalkowski/go-service/transport"
	"go.uber.org/fx"
)

// ServerOptions for cmd.
var ServerOptions = []fx.Option{
	fx.NopLogger, runtime.Module, debug.Module,
	telemetry.Module, metrics.Module, Module,
	config.Module, health.Module, transport.Module,
}
