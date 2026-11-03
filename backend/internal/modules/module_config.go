package modules

import (
	"github.com/kyawzinthant-coding/froky/config"
	"go.uber.org/fx"
)

var Config = fx.Module("configuration",
	fx.Provide(
		config.LoadConfig,
	),
)
