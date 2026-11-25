package vendor_menu

import (
	"context"
	"testing"

	"github.com/kyawzinthant-coding/froky/api"
	"github.com/kyawzinthant-coding/froky/internal/service/menu"
	"github.com/kyawzinthant-coding/froky/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func productBody() *api.ProductRequest {
	return &api.ProductRequest{Name: "Pad Thai", Description: "Noodles", Price: 6500, CategoryId: 1, IsAvailable: true}
}

func TestProducts_NoClaims_Unauthorized(t *testing.T) {
	t.Parallel()

	fake := &fakeMenu{}
	h := &VendorMenuHandler{menuService: fake}

	resp, err := h.CreateVendorProduct(context.Background(), api.CreateVendorProductRequestObject{Body: productBody()})
	require.NoError(t, err)
	assert.IsType(t, api.CreateVendorProduct401JSONResponse{}, resp)
	assert.Zero(t, fake.gotVendorID, "service must not be called without a session")
}

func TestCreateProduct(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		fake *fakeMenu
		want api.CreateVendorProductResponseObject
	}{
		{
			name: "created",
			fake: &fakeMenu{product: &types.Product{ID: 3, Name: "Pad Thai", Price: 6500, CategoryID: 1, IsAvailable: true}},
			want: api.CreateVendorProduct201JSONResponse{Id: 3, Name: "Pad Thai", Price: 6500, CategoryId: 1, IsAvailable: true},
		},
		{
			name: "invalid input",
			fake: &fakeMenu{err: menu.ErrInvalidInput},
			want: api.CreateVendorProduct400JSONResponse{BadRequestJSONResponse: api.BadRequestJSONResponse{Error: menu.ErrInvalidInput.Error()}},
		},
		{
			name: "category of another vendor",
			fake: &fakeMenu{err: menu.ErrCategoryNotFound},
			want: api.CreateVendorProduct404JSONResponse{Error: msgCategoryNotFound},
		},
		{
			name: "name taken",
			fake: &fakeMenu{err: menu.ErrNameTaken},
			want: api.CreateVendorProduct409JSONResponse{ProductNameTakenJSONResponse: api.ProductNameTakenJSONResponse{Error: msgProductNameTaken}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := &VendorMenuHandler{menuService: tc.fake}

			resp, err := h.CreateVendorProduct(loggedInCtx(), api.CreateVendorProductRequestObject{Body: productBody()})
			require.NoError(t, err)
			assert.Equal(t, tc.want, resp)
			assert.Equal(t, testVendorID, tc.fake.gotVendorID, "vendor ID must come from the session")
		})
	}
}

func TestUpdateProduct_NotFound(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		err     error
		wantMsg string
	}{
		{name: "product", err: menu.ErrProductNotFound, wantMsg: msgProductNotFound},
		{name: "category", err: menu.ErrCategoryNotFound, wantMsg: msgCategoryNotFound},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := &VendorMenuHandler{menuService: &fakeMenu{err: tc.err}}

			resp, err := h.UpdateVendorProduct(loggedInCtx(), api.UpdateVendorProductRequestObject{ProductId: 3, Body: productBody()})
			require.NoError(t, err)
			assert.Equal(t, api.UpdateVendorProduct404JSONResponse{Error: tc.wantMsg}, resp)
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	t.Parallel()

	h := &VendorMenuHandler{menuService: &fakeMenu{err: menu.ErrProductNotFound}}
	resp, err := h.DeleteVendorProduct(loggedInCtx(), api.DeleteVendorProductRequestObject{ProductId: 3})
	require.NoError(t, err)
	assert.IsType(t, api.DeleteVendorProduct404JSONResponse{}, resp)

	h = &VendorMenuHandler{menuService: &fakeMenu{}}
	resp, err = h.DeleteVendorProduct(loggedInCtx(), api.DeleteVendorProductRequestObject{ProductId: 3})
	require.NoError(t, err)
	assert.Equal(t, api.DeleteVendorProduct204Response{}, resp)
}
