package vendor_auth

import (
	"context"
	"errors"
	"github.com/kyawzinthant-coding/froky/api"
	"github.com/kyawzinthant-coding/froky/internal/service/auth"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

func (h *VendorAuthHandler) VendorLogin(ctx context.Context, request api.VendorLoginRequestObject) (api.VendorLoginResponseObject, error) {
	username := strings.TrimSpace(request.Body.Username)
	password := request.Body.Password

	if username == "" || password == "" {
		return api.VendorLogin401JSONResponse{
			Error: "invalid credentials",
		}, nil
	}

	token, vendor, err := h.authService.Login(ctx, username, password)
	if err != nil {
		if errors.Is(err, auth.ErrInvalidCredentials) {
			log.Warn().Str("username", username).Msg("vendor login denied: invalid credentials")
			return api.VendorLogin401JSONResponse{Error: "invalid credentials"}, nil
		}
		log.Error().Err(err).Str("username", username).Msg("vendor login failed: unexpected error")
		return nil, err
	}

	cookie := (&http.Cookie{
		Name:     vendorCookieName,
		Value:    token,
		Path:     "/",
		MaxAge:   int(h.tokenTTL.Seconds()),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}).String()

	return api.VendorLogin200JSONResponse{
		Body: api.VendorResponse{
			Id:       vendor.ID,
			Name:     vendor.Name,
			Username: vendor.Username,
		},
		Headers: api.VendorLogin200ResponseHeaders{SetCookie: &cookie},
	}, nil
}
