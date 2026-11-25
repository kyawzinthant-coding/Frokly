package public_menu

import (
	"context"
	"errors"
	"testing"

	"github.com/kyawzinthant-coding/froky/api"
	"github.com/kyawzinthant-coding/froky/internal/service/vendor"
	"github.com/kyawzinthant-coding/froky/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeVendors struct {
	vendors []types.Vendor
	vendor  *types.Vendor
	err     error
}

func (f *fakeVendors) List(context.Context) ([]types.Vendor, error) { return f.vendors, f.err }

func (f *fakeVendors) Get(context.Context, int64) (*types.Vendor, error) { return f.vendor, f.err }

type fakeMenu struct {
	sections []types.MenuSection
	err      error
	called   bool
}

func (f *fakeMenu) AvailableMenu(context.Context, int64) ([]types.MenuSection, error) {
	f.called = true
	return f.sections, f.err
}

func strPtr(s string) *string { return &s }

func TestListPublicVendors(t *testing.T) {
	t.Parallel()

	h := &PublicMenuHandler{
		vendorService: &fakeVendors{vendors: []types.Vendor{
			{ID: 1, Name: "Khun Mae Noodles", LogoURL: strPtr("https://cdn.test/khunmae.png")},
			{ID: 2, Name: "Somtam Corner"},
		}},
		menuService: &fakeMenu{},
	}

	resp, err := h.ListPublicVendors(context.Background(), api.ListPublicVendorsRequestObject{})
	require.NoError(t, err)
	assert.Equal(t, api.ListPublicVendors200JSONResponse{
		{Id: 1, Name: "Khun Mae Noodles", LogoUrl: strPtr("https://cdn.test/khunmae.png")},
		{Id: 2, Name: "Somtam Corner"},
	}, resp)
}

func TestListPublicVendors_EmptyIsArray(t *testing.T) {
	t.Parallel()

	h := &PublicMenuHandler{vendorService: &fakeVendors{}, menuService: &fakeMenu{}}

	resp, err := h.ListPublicVendors(context.Background(), api.ListPublicVendorsRequestObject{})
	require.NoError(t, err)
	assert.NotNil(t, resp, "an empty list must encode as [] not null")
	assert.Empty(t, resp)
}

func TestGetPublicVendorMenu(t *testing.T) {
	t.Parallel()

	menus := &fakeMenu{sections: []types.MenuSection{{
		ID:   10,
		Name: "Noodle Soups",
		Items: []types.Product{
			{ID: 100, Name: "Boat Noodles", Description: "Pork broth", Price: 6000, CategoryID: 10, IsAvailable: true},
		},
	}}}
	h := &PublicMenuHandler{
		vendorService: &fakeVendors{vendor: &types.Vendor{ID: 1, Name: "Khun Mae Noodles"}},
		menuService:   menus,
	}

	resp, err := h.GetPublicVendorMenu(context.Background(), api.GetPublicVendorMenuRequestObject{VendorId: 1})
	require.NoError(t, err)
	assert.Equal(t, api.GetPublicVendorMenu200JSONResponse{
		Vendor: api.PublicVendor{Id: 1, Name: "Khun Mae Noodles"},
		Categories: []api.PublicMenuCategory{{
			Id:    10,
			Name:  "Noodle Soups",
			Items: []api.PublicMenuItem{{Id: 100, Name: "Boat Noodles", Description: "Pork broth", Price: 6000}},
		}},
	}, resp)
}

func TestGetPublicVendorMenu_UnknownVendor(t *testing.T) {
	t.Parallel()

	menus := &fakeMenu{}
	h := &PublicMenuHandler{vendorService: &fakeVendors{err: vendor.ErrVendorNotFound}, menuService: menus}

	resp, err := h.GetPublicVendorMenu(context.Background(), api.GetPublicVendorMenuRequestObject{VendorId: 999})
	require.NoError(t, err)
	assert.IsType(t, api.GetPublicVendorMenu404JSONResponse{}, resp)
	assert.False(t, menus.called, "no menu lookup for a stall that doesn't exist")
}

func TestGetPublicVendorMenu_UnexpectedError(t *testing.T) {
	t.Parallel()

	h := &PublicMenuHandler{
		vendorService: &fakeVendors{vendor: &types.Vendor{ID: 1, Name: "Khun Mae Noodles"}},
		menuService:   &fakeMenu{err: errors.New("db down")},
	}

	resp, err := h.GetPublicVendorMenu(context.Background(), api.GetPublicVendorMenuRequestObject{VendorId: 1})
	require.Error(t, err)
	assert.Nil(t, resp)
}
