package modules

import (
	"github.com/kyawzinthant-coding/froky/config"
	"go.uber.org/fx"
)

var Config = fx.Module("configuration",
	fx.Provide(
		config.LoadConfig,
	),

	fx.Provide(func(cfg *config.AppConfig) *config.DatabaseConfig { return cfg.DataBase }),
)
