package vendor_menu

import (
	"context"
	"errors"
	"testing"

	"github.com/kyawzinthant-coding/froky/api"
	"github.com/kyawzinthant-coding/froky/internal/service/menu"
	"github.com/kyawzinthant-coding/froky/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCategories_NoClaims_Unauthorized(t *testing.T) {
	t.Parallel()

	fake := &fakeMenu{}
	h := &VendorMenuHandler{menuService: fake}
	ctx := context.Background()

	list, err := h.ListVendorCategories(ctx, api.ListVendorCategoriesRequestObject{})
	require.NoError(t, err)
	assert.IsType(t, api.ListVendorCategories401JSONResponse{}, list)

	del, err := h.DeleteVendorCategory(ctx, api.DeleteVendorCategoryRequestObject{CategoryId: 1})
	require.NoError(t, err)
	assert.IsType(t, api.DeleteVendorCategory401JSONResponse{}, del)

	assert.Zero(t, fake.gotVendorID, "service must not be called without a session")
}

func TestCreateCategory(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		fake *fakeMenu
		want api.CreateVendorCategoryResponseObject
	}{
		{
			name: "created",
			fake: &fakeMenu{category: &types.Category{ID: 1, Name: "Noodles"}},
			want: api.CreateVendorCategory201JSONResponse{Id: 1, Name: "Noodles"},
		},
		{
			name: "invalid input",
			fake: &fakeMenu{err: menu.ErrInvalidInput},
			want: api.CreateVendorCategory400JSONResponse{BadRequestJSONResponse: api.BadRequestJSONResponse{Error: menu.ErrInvalidInput.Error()}},
		},
		{
			name: "name taken",
			fake: &fakeMenu{err: menu.ErrNameTaken},
			want: api.CreateVendorCategory409JSONResponse{CategoryNameTakenJSONResponse: api.CategoryNameTakenJSONResponse{Error: msgCategoryNameTaken}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := &VendorMenuHandler{menuService: tc.fake}

			resp, err := h.CreateVendorCategory(loggedInCtx(), api.CreateVendorCategoryRequestObject{
				Body: &api.CreateVendorCategoryJSONRequestBody{Name: "Noodles"},
			})
			require.NoError(t, err)
			assert.Equal(t, tc.want, resp)
			assert.Equal(t, testVendorID, tc.fake.gotVendorID, "vendor ID must come from the session")
		})
	}
}

func TestUpdateCategory_NotFound(t *testing.T) {
	t.Parallel()

	h := &VendorMenuHandler{menuService: &fakeMenu{err: menu.ErrCategoryNotFound}}

	resp, err := h.UpdateVendorCategory(loggedInCtx(), api.UpdateVendorCategoryRequestObject{
		CategoryId: 99,
		Body:       &api.UpdateVendorCategoryJSONRequestBody{Name: "Noodles"},
	})
	require.NoError(t, err)
	assert.IsType(t, api.UpdateVendorCategory404JSONResponse{}, resp)
}

func TestDeleteCategory(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		want api.DeleteVendorCategoryResponseObject
	}{
		{name: "archived", err: nil, want: api.DeleteVendorCategory204Response{}},
		{
			name: "not found",
			err:  menu.ErrCategoryNotFound,
			want: api.DeleteVendorCategory404JSONResponse{CategoryNotFoundJSONResponse: api.CategoryNotFoundJSONResponse{Error: msgCategoryNotFound}},
		},
		{
			name: "not empty",
			err:  menu.ErrCategoryNotEmpty,
			want: api.DeleteVendorCategory409JSONResponse{CategoryNotEmptyJSONResponse: api.CategoryNotEmptyJSONResponse{Error: msgCategoryNotEmpty}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := &VendorMenuHandler{menuService: &fakeMenu{err: tc.err}}

			resp, err := h.DeleteVendorCategory(loggedInCtx(), api.DeleteVendorCategoryRequestObject{CategoryId: 1})
			require.NoError(t, err)
			assert.Equal(t, tc.want, resp)
		})
	}
}

func TestListCategories_UnexpectedError(t *testing.T) {
	t.Parallel()

	h := &VendorMenuHandler{menuService: &fakeMenu{err: errors.New("db down")}}

	resp, err := h.ListVendorCategories(loggedInCtx(), api.ListVendorCategoriesRequestObject{})
	require.Error(t, err)
	assert.Nil(t, resp)
}
