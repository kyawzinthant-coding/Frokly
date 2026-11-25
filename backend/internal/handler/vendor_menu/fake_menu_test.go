package vendor_menu

import (
	"context"

	"github.com/kyawzinthant-coding/froky/internal/middleware"
	"github.com/kyawzinthant-coding/froky/internal/types"
)

const testVendorID int64 = 7

// fakeMenu returns err from every method and records the vendor ID it was called with.
type fakeMenu struct {
	err         error
	category    *types.Category
	product     *types.Product
	gotVendorID int64
}

func (f *fakeMenu) ListCategories(_ context.Context, vendorID int64) ([]types.Category, error) {
	f.gotVendorID = vendorID
	return nil, f.err
}

func (f *fakeMenu) GetCategory(_ context.Context, vendorID, _ int64) (*types.Category, error) {
	f.gotVendorID = vendorID
	return f.category, f.err
}

func (f *fakeMenu) CreateCategory(_ context.Context, vendorID int64, _ string) (*types.Category, error) {
	f.gotVendorID = vendorID
	return f.category, f.err
}

func (f *fakeMenu) UpdateCategory(_ context.Context, vendorID, _ int64, _ string) (*types.Category, error) {
	f.gotVendorID = vendorID
	return f.category, f.err
}

func (f *fakeMenu) ArchiveCategory(_ context.Context, vendorID, _ int64) error {
	f.gotVendorID = vendorID
	return f.err
}

func (f *fakeMenu) ListProducts(_ context.Context, vendorID int64) ([]types.Product, error) {
	f.gotVendorID = vendorID
	return nil, f.err
}

func (f *fakeMenu) GetProduct(_ context.Context, vendorID, _ int64) (*types.Product, error) {
	f.gotVendorID = vendorID
	return f.product, f.err
}

func (f *fakeMenu) CreateProduct(_ context.Context, vendorID int64, _ types.ProductInput) (*types.Product, error) {
	f.gotVendorID = vendorID
	return f.product, f.err
}

func (f *fakeMenu) UpdateProduct(_ context.Context, vendorID, _ int64, _ types.ProductInput) (*types.Product, error) {
	f.gotVendorID = vendorID
	return f.product, f.err
}

func (f *fakeMenu) ArchiveProduct(_ context.Context, vendorID, _ int64) error {
	f.gotVendorID = vendorID
	return f.err
}

func loggedInCtx() context.Context {
	return middleware.ContextWithVendorClaims(context.Background(), &types.VendorClaims{VendorID: testVendorID})
}
