package vendor_auth

import (
	"context"
	"github.com/kyawzinthant-coding/froky/api"
)

func (h *VendorAuthHandler) VendorMe(ctx context.Context, request api.VendorMeRequestObject) (api.VendorMeResponseObject, error) {
	return api.VendorMe401JSONResponse{Error: "not implemented"}, nil
}
