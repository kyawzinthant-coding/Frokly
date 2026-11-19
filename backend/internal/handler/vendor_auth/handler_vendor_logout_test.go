package vendor_auth

import (
	"context"
	"github.com/kyawzinthant-coding/froky/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVendorLogout_ClearsCookie(t *testing.T) {
	t.Parallel()

	h := newTestHandler(&fakeAuth{})
	resp, err := h.VendorLogout(context.Background(), api.VendorLogoutRequestObject{})
	require.NoError(t, err)

	rec := httptest.NewRecorder()
	require.NoError(t, resp.VisitVendorLogoutResponse(rec))

	assert.Equal(t, http.StatusNoContent, rec.Code)

	cookie := rec.Header().Get("Set-Cookie")
	assert.Contains(t, cookie, "vendor_token=;")
	assert.Contains(t, cookie, "Path=/")
	assert.Contains(t, cookie, "Max-Age=0")
	assert.Contains(t, cookie, "HttpOnly")
}
