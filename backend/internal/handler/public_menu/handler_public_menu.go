package public_menu

import (
	"context"
	"errors"

	"github.com/kyawzinthant-coding/froky/api"
	"github.com/kyawzinthant-coding/froky/internal/service/vendor"
	"github.com/rs/zerolog/log"
)

func (h *PublicMenuHandler) ListPublicVendors(ctx context.Context, _ api.ListPublicVendorsRequestObject) (api.ListPublicVendorsResponseObject, error) {
	vendors, err := h.vendorService.List(ctx)
	if err != nil {
		log.Error().Err(err).Msg("list public vendors failed")
		return nil, err
	}

	resp := make(api.ListPublicVendors200JSONResponse, 0, len(vendors))
	for _, v := range vendors {
		resp = append(resp, toPublicVendor(v))
	}
	return resp, nil
}

// The vendor is looked up first so an unknown stall is a 404, not an empty menu.
func (h *PublicMenuHandler) GetPublicVendorMenu(ctx context.Context, req api.GetPublicVendorMenuRequestObject) (api.GetPublicVendorMenuResponseObject, error) {
	v, err := h.vendorService.Get(ctx, req.VendorId)
	switch {
	case errors.Is(err, vendor.ErrVendorNotFound):
		return api.GetPublicVendorMenu404JSONResponse{VendorNotFoundJSONResponse: api.VendorNotFoundJSONResponse{Error: msgVendorNotFound}}, nil
	case err != nil:
		log.Error().Err(err).Int64("vendor_id", req.VendorId).Msg("get public vendor failed")
		return nil, err
	}

	sections, err := h.menuService.AvailableMenu(ctx, v.ID)
	if err != nil {
		log.Error().Err(err).Int64("vendor_id", v.ID).Msg("get public menu failed")
		return nil, err
	}

	categories := make([]api.PublicMenuCategory, 0, len(sections))
	for _, section := range sections {
		categories = append(categories, toPublicCategory(section))
	}

	return api.GetPublicVendorMenu200JSONResponse{Vendor: toPublicVendor(*v), Categories: categories}, nil
}
