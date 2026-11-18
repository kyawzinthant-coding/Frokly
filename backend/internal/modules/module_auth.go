package modules

import (
	"github.com/kyawzinthant-coding/froky/internal/service/auth"
	"go.uber.org/fx"
)

var AuthModule = fx.Module("auth",
	fx.Provide(auth.NewService),
)
