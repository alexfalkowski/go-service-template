package health

import (
	"github.com/alexfalkowski/go-service-template/internal/health/transport/grpc"
	"github.com/alexfalkowski/go-service-template/internal/health/transport/http"
	"github.com/alexfalkowski/go-service/v2/health"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	fx.Provide(http.NewHealthObserver), fx.Provide(http.NewLivenessObserver), fx.Provide(http.NewReadinessObserver),
	fx.Provide(grpc.NewObserver), fx.Provide(NewRegistrations), health.Module,
)
