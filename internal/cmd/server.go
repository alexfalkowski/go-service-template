package cmd

import (
	"github.com/alexfalkowski/go-service-template/internal/config"
	"github.com/alexfalkowski/go-service-template/internal/health"
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/module"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/transport"
)

// RegisterServer for cmd.
func RegisterServer(command *cmd.Command) {
	flags := command.AddServer("server", "Start go-service-template server",
		module.Module, debug.Module, feature.Module,
		telemetry.Module, transport.Module,
		config.Module, health.Module, cmd.Module,
	)
	flags.AddInput("")
}
