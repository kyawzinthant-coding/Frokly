package modules

import (
	"github.com/kyawzinthant-coding/froky/internal/service/menu"
	"go.uber.org/fx"
)

var MenuModule = fx.Module("menu",
	fx.Provide(menu.NewService),
)
