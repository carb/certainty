package controllers

import (
	"go.uber.org/fx"
)

// Module returns an fx.Option that provides an contructors to an Fx application.
var Module = fx.Option(
	fx.Provide(
		NewBoardController,
	),
)
