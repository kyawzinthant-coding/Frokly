package handler

import (
	"database/sql"
	"github.com/kyawzinthant-coding/froky/api"
	"github.com/kyawzinthant-coding/froky/config"
	"github.com/kyawzinthant-coding/froky/internal/handler/vendor_auth"
	"github.com/kyawzinthant-coding/froky/internal/service/auth"
)

var _ api.StrictServerInterface = (*Handler)(nil)

type Handler struct {
	*vendor_auth.VendorAuthHandler
	db *sql.DB
}

func NewHandler(
	cfg *config.AppConfig,
	vendorService *auth.Service,
	db *sql.DB,
) *Handler {
	return &Handler{
		db:                db,
		VendorAuthHandler: vendor_auth.NewVendorAuthHandler(cfg, vendorService),
	}
}
