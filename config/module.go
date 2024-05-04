package config

import (
	"github.com/alexfalkowski/go-service/config"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	config.Module,
	fx.Provide(NewConfig),
	fx.Decorate(decorateConfig),
	fx.Provide(healthConfig),
)
