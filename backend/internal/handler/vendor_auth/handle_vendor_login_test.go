package vendor_auth

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/kyawzinthant-coding/froky/api"
	"github.com/kyawzinthant-coding/froky/internal/service/auth"
	"github.com/kyawzinthant-coding/froky/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type fakeAuth struct {
	token   string
	profile *types.VendorProfile
	err     error

	called      bool
	gotUsername string
	gotPassword string
}

func (f *fakeAuth) Login(_ context.Context, username, password string) (string, *types.VendorProfile, error) {
	f.called = true
	f.gotUsername = username
	f.gotPassword = password

	return f.token, f.profile, f.err
}

func (f *fakeAuth) Me(_ context.Context, _ int64) (*types.VendorProfile, error) {
	return f.profile, f.err
}

func newTestHandler(fake *fakeAuth) *VendorAuthHandler {
	return &VendorAuthHandler{
		authService: fake,
		tokenTTL:    time.Hour,
	}
}

func loginRequest(username, password string) api.VendorLoginRequestObject {
	return api.VendorLoginRequestObject{
		Body: &api.VendorLoginJSONRequestBody{Username: username, Password: password},
	}
}

func TestVendorLogin_Success(t *testing.T) {
	t.Parallel()

	fake := &fakeAuth{
		token: "tok-123",
		profile: &types.VendorProfile{
			ID:       1,
			Name:     "Test Vendor",
			Username: "vendor1",
		},
	}

	h := newTestHandler(fake)
	resp, err := h.VendorLogin(context.Background(), loginRequest("  vendor1  ", "password123"))
	require.NoError(t, err)

	assert.Equal(t, "vendor1", fake.gotUsername)
	assert.Equal(t, "password123", fake.gotPassword)

	rec := httptest.NewRecorder()
	require.NoError(t, resp.VisitVendorLoginResponse(rec))

	assert.Equal(t, http.StatusOK, rec.Code)

	cookie := rec.Header().Get("Set-Cookie")
	assert.Contains(t, cookie, "vendor_token=tok-123")
	assert.Contains(t, cookie, "HttpOnly")
	assert.Contains(t, cookie, "Max-Age=3600")

	var body api.VendorResponse
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&body))
	assert.Equal(t, api.VendorResponse{Id: 1, Name: "Test Vendor", Username: "vendor1"}, body)
}

func TestVendorLogin_Rejected(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		username   string
		password   string
		fake       *fakeAuth
		wantCalled bool
	}{
		{name: "empty username", username: "", password: "pw", fake: &fakeAuth{}, wantCalled: false},
		{name: "whitespace-only username", username: "   ", password: "pw", fake: &fakeAuth{}, wantCalled: false},
		{name: "empty password", username: "vendor1", password: "", fake: &fakeAuth{}, wantCalled: false},
		{
			name:       "invalid credentials",
			username:   "vendor1",
			password:   "wrong",
			fake:       &fakeAuth{err: auth.ErrInvalidCredentials},
			wantCalled: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := newTestHandler(tc.fake)

			resp, err := h.VendorLogin(context.Background(), loginRequest(tc.username, tc.password))
			require.NoError(t, err)

			require.IsType(t, api.VendorLogin401JSONResponse{}, resp)
			assert.Equal(t, "invalid credentials", resp.(api.VendorLogin401JSONResponse).Error)
			assert.Equal(t, tc.wantCalled, tc.fake.called)
		})
	}
}

func TestVendorLogin_UnexpectedError(t *testing.T) {
	t.Parallel()

	fake := &fakeAuth{err: errors.New("db down")}
	h := newTestHandler(fake)

	resp, err := h.VendorLogin(context.Background(), loginRequest("vendor1", "password123"))

	require.Error(t, err)
	assert.Nil(t, resp)
}
