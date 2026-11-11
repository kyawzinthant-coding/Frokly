package modules

import (
	"github.com/kyawzinthant-coding/froky/internal/handler"
	"go.uber.org/fx"
)

var HandlerModule = fx.Module("handler",
	fx.Provide(handler.NewHandler),
)
