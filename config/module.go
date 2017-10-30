package config

import (
	"go.uber.org/fx"
)

// Module returns an fx.Option that provides an *Configuration
// into an Fx application.
var Module = fx.Option(
	fx.Provide(
		GetConfiguration,
		NewSugaredLogger,
	),
)
