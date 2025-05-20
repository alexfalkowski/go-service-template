package cmd

import (
	"github.com/alexfalkowski/go-service-template/internal/config"
	"github.com/alexfalkowski/go-service-template/internal/health"
	"github.com/alexfalkowski/go-service/v2/cli"
	"github.com/alexfalkowski/go-service/v2/debug"
	"github.com/alexfalkowski/go-service/v2/feature"
	"github.com/alexfalkowski/go-service/v2/module"
	"github.com/alexfalkowski/go-service/v2/telemetry"
	"github.com/alexfalkowski/go-service/v2/transport"
)

// RegisterServer for cmd.
func RegisterServer(command cli.Commander) {
	cmd := command.AddServer("server", "Start go-service-template server",
		module.Module, debug.Module, feature.Module,
		telemetry.Module, transport.Module,
		config.Module, health.Module, cli.Module,
	)
	cmd.AddInput("")
}
