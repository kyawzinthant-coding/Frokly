package vendor_auth

import (
	"context"
	"github.com/kyawzinthant-coding/froky/api"
)

func (h *VendorAuthHandler) VendorLogout(_ context.Context, _ api.VendorLogoutRequestObject) (api.VendorLogoutResponseObject, error) {
	cookie := vendorCookie("", deleteCookieMaxAge)

	return api.VendorLogout204Response{
		Headers: api.VendorLogout204ResponseHeaders{SetCookie: &cookie},
	}, nil
}
