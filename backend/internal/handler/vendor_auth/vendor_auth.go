package vendor_auth

import (
	"context"
	"github.com/kyawzinthant-coding/froky/config"
	"github.com/kyawzinthant-coding/froky/internal/service/auth"
	"github.com/kyawzinthant-coding/froky/internal/types"
	"net/http"
	"time"
)

// Go's http.Cookie treats MaxAge 0 as "no Max-Age"; a negative value sends Max-Age=0, which deletes it.
const deleteCookieMaxAge = -1

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

// Login and logout must send identical cookie attributes, or the browser won't replace the cookie.
func vendorCookie(value string, maxAge int) string {
	return (&http.Cookie{
		Name:     types.VendorCookiesName,
		Value:    value,
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}).String()
}
