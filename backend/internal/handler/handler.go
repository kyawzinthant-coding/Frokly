package handler

import (
	"database/sql"
	"github.com/kyawzinthant-coding/froky/api"
	"github.com/kyawzinthant-coding/froky/config"
	"github.com/kyawzinthant-coding/froky/internal/handler/public_menu"
	"github.com/kyawzinthant-coding/froky/internal/handler/vendor_auth"
	"github.com/kyawzinthant-coding/froky/internal/handler/vendor_menu"
	"github.com/kyawzinthant-coding/froky/internal/service/auth"
	"github.com/kyawzinthant-coding/froky/internal/service/menu"
	"github.com/kyawzinthant-coding/froky/internal/service/vendor"
)

var _ api.StrictServerInterface = (*Handler)(nil)

type Handler struct {
	*vendor_auth.VendorAuthHandler
	*vendor_menu.VendorMenuHandler
	*public_menu.PublicMenuHandler
	db *sql.DB
}

func NewHandler(
	cfg *config.AppConfig,
	vendorService *auth.Service,
	menuService *menu.Service,
	vendors *vendor.Service,
	db *sql.DB,
) *Handler {
	return &Handler{
		db:                db,
		VendorAuthHandler: vendor_auth.NewVendorAuthHandler(cfg, vendorService),
		VendorMenuHandler: vendor_menu.NewVendorMenuHandler(menuService),
		PublicMenuHandler: public_menu.NewPublicMenuHandler(vendors, menuService),
	}
}
