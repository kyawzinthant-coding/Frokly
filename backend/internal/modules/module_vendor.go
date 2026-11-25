package modules

import (
	"github.com/kyawzinthant-coding/froky/internal/service/vendor"
	"go.uber.org/fx"
)

var VendorModule = fx.Module("vendor",
	fx.Provide(vendor.NewService),
)
