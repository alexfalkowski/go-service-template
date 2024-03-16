package health

import (
	"time"

	"github.com/alexfalkowski/go-health/checker"
	"github.com/alexfalkowski/go-health/server"
	h "github.com/alexfalkowski/go-service-template/health"
	"github.com/alexfalkowski/go-service/health"
	"go.uber.org/fx"
)

// Params for health.
type Params struct {
	fx.In

	Health *h.Config
}

// NewRegistrations for health.
func NewRegistrations(params Params) (health.Registrations, error) {
	d, err := time.ParseDuration(params.Health.Duration)
	if err != nil {
		return nil, err
	}

	registrations := health.Registrations{
		server.NewRegistration("noop", d, checker.NewNoopChecker()),
	}

	return registrations, nil
}
