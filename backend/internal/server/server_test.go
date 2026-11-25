package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/kyawzinthant-coding/froky/api"
	"github.com/stretchr/testify/assert"
)

// Runs the real generated router with handlers that answer 501, so a request that gets
// past auth shows 501 and one stopped by auth shows 401. No database or services needed.
func TestRequireVendorAuth_FollowsSpecSecurity(t *testing.T) {
	t.Parallel()

	router := api.HandlerWithOptions(api.Unimplemented{}, api.ChiServerOptions{
		BaseRouter:  chi.NewRouter(),
		Middlewares: []api.MiddlewareFunc{requireVendorAuth([]byte("test-secret"))},
	})

	tests := []struct {
		method, path string
		want         int
	}{
		{http.MethodGet, "/v1/status", http.StatusNotImplemented},
		{http.MethodPost, "/v1/vendor/login", http.StatusNotImplemented},
		{http.MethodPost, "/v1/vendor/logout", http.StatusNotImplemented},
		{http.MethodGet, "/v1/vendors", http.StatusNotImplemented},
		{http.MethodGet, "/v1/vendors/1/menu", http.StatusNotImplemented},
		{http.MethodGet, "/v1/vendor/me", http.StatusUnauthorized},
		{http.MethodGet, "/v1/vendor/products", http.StatusUnauthorized},
		{http.MethodGet, "/v1/vendor/products/1", http.StatusUnauthorized},
		{http.MethodDelete, "/v1/vendor/categories/1", http.StatusUnauthorized},
	}

	for _, tc := range tests {
		t.Run(tc.method+" "+tc.path, func(t *testing.T) {
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, httptest.NewRequest(tc.method, tc.path, nil))
			assert.Equal(t, tc.want, rec.Code)
		})
	}
}
