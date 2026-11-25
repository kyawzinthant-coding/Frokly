package handler

import (
	"database/sql"
	"github.com/kyawzinthant-coding/froky/api"
	"github.com/kyawzinthant-coding/froky/config"
	"github.com/kyawzinthant-coding/froky/internal/handler/vendor_auth"
	"github.com/kyawzinthant-coding/froky/internal/handler/vendor_menu"
	"github.com/kyawzinthant-coding/froky/internal/service/auth"
	"github.com/kyawzinthant-coding/froky/internal/service/menu"
)

var _ api.StrictServerInterface = (*Handler)(nil)

type Handler struct {
	*vendor_auth.VendorAuthHandler
	*vendor_menu.VendorMenuHandler
	db *sql.DB
}

func NewHandler(
	cfg *config.AppConfig,
	vendorService *auth.Service,
	menuService *menu.Service,
	db *sql.DB,
) *Handler {
	return &Handler{
		db:                db,
		VendorAuthHandler: vendor_auth.NewVendorAuthHandler(cfg, vendorService),
		VendorMenuHandler: vendor_menu.NewVendorMenuHandler(menuService),
	}
}
