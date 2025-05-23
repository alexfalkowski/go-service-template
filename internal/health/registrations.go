package health

import (
	"github.com/alexfalkowski/go-health/checker"
	"github.com/alexfalkowski/go-health/server"
	"github.com/alexfalkowski/go-service/v2/health"
	"github.com/alexfalkowski/go-service/v2/time"
	"go.uber.org/fx"
)

// Params for health.
type Params struct {
	fx.In

	Health *Config
}

// NewRegistrations for health.
func NewRegistrations(params Params) health.Registrations {
	d := time.MustParseDuration(params.Health.Duration)
	registrations := health.Registrations{
		server.NewRegistration("noop", d, checker.NewNoopChecker()),
		server.NewOnlineRegistration(d, d),
	}

	return registrations
}
