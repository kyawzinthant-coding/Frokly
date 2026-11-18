package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kyawzinthant-coding/froky/internal/types"
	"github.com/rs/zerolog/log"
	"net/http"
)

type vendorClaimsKey struct{}

func VendorClaimsFromContext(ctx context.Context) (*types.VendorClaims, bool) {
	claims, ok := ctx.Value(vendorClaimsKey{}).(*types.VendorClaims)
	return claims, ok
}

func VendorJwtAuth(secret []byte) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			cookie, err := r.Cookie(types.VendorCookiesName)
			if errors.Is(err, http.ErrNoCookie) || (err == nil && cookie.Value == "") {
				log.Debug().Str("path", r.URL.Path).Str("method", r.Method).Msg("no cookie found")
				writeUnauthorized(w, "unauthorized")
				return
			}

			if err != nil {
				log.Warn().Err(err).Str("path", r.URL.Path).Msg("vendor auth: bad cookie")
				writeUnauthorized(w, "unauthorized")
				return
			}

			tokenStr := cookie.Value
			claims := &types.VendorClaims{}
			token, err := jwt.ParseWithClaims(tokenStr, claims,
				func(token *jwt.Token) (interface{}, error) {
					return secret, nil
				},
				jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
			)

			if err != nil || !token.Valid {
				log.Warn().Err(err).Msg("vendor auth: bad token")
				writeUnauthorized(w, "unauthorized")
				return
			}

			log.Debug().
				Msg("vendor auth: valid")

			ctx := context.WithValue(r.Context(), vendorClaimsKey{}, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func writeUnauthorized(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
