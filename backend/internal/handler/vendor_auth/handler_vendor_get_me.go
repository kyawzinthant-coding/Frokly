package vendor_auth

import (
	"context"
	"errors"
	"github.com/kyawzinthant-coding/froky/api"
	"github.com/kyawzinthant-coding/froky/internal/middleware"
	"github.com/kyawzinthant-coding/froky/internal/service/auth"
	"github.com/rs/zerolog/log"
	"strconv"
)

func (h *VendorAuthHandler) VendorMe(ctx context.Context, request api.VendorMeRequestObject) (api.VendorMeResponseObject, error) {
	claims, ok := middleware.VendorClaimsFromContext(ctx)
	if !ok {
		log.Error().Msg("vendor me: no claims in context (is the route behind the auth middleware?)")
		return api.VendorMe401JSONResponse{Error: "unauthorized"}, nil
	}

	staffID, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		log.Warn().Str("subject", claims.Subject).Msg("invalid subject")
		return api.VendorMe401JSONResponse{Error: "unauthorized"}, nil
	}

	profile, err := h.authService.Me(ctx, staffID)

	if err != nil {
		if errors.Is(err, auth.ErrStaffNotFound) {
			log.Warn().Int64("staff_id", staffID).Msg("vendor me: staff no longer exists")
			return api.VendorMe401JSONResponse{Error: "unauthorized"}, nil
		}
		log.Error().Err(err).Int64("staff_id", staffID).Msg("vendor me: unexpected error")
		return nil, err
	}

	return api.VendorMe200JSONResponse{
		Id:       profile.ID,
		Name:     profile.Name,
		Username: profile.Username,
	}, nil
}
