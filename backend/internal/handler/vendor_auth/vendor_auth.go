package vendor_auth

import (
	"context"
	"github.com/kyawzinthant-coding/froky/config"
	"github.com/kyawzinthant-coding/froky/internal/service/auth"
	"github.com/kyawzinthant-coding/froky/internal/types"
	"time"
)

type authService interface {
	Login(ctx context.Context, username, password string) (string, *types.VendorProfile, error)
	Me(ctx context.Context, staffId int64) (*types.VendorProfile, error)
}

type VendorAuthHandler struct {
	authService authService
	tokenTTL    time.Duration
}

func NewVendorAuthHandler(cfg *config.AppConfig, svc *auth.Service) *VendorAuthHandler {
	return &VendorAuthHandler{
		authService: svc,
		tokenTTL:    cfg.JWT.Expiry,
	}
}
