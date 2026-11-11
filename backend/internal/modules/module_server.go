package modules

import (
	"github.com/kyawzinthant-coding/froky/internal/server"
	"go.uber.org/fx"
)

var ServerModule = fx.Module("server",
	fx.Provide(server.NewHttpServer),
	fx.Invoke((*server.HttpServer).Start),
)
