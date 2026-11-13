package cmd

import (
	"github.com/alexfalkowski/go-service-template/v2/internal/config"
	"github.com/alexfalkowski/go-service-template/v2/internal/health"
	"github.com/alexfalkowski/go-service/v3/module"
	"github.com/alexfalkowski/go-signal/v2"
)

// Module composes the server with service configuration and health checks.
var Module = signal.Module(
	module.Server,
	config.Module,
	health.Module,
)
