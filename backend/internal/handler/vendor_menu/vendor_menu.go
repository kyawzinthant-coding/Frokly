package vendor_menu

import (
	"context"

	"github.com/kyawzinthant-coding/froky/api"
	"github.com/kyawzinthant-coding/froky/internal/middleware"
	"github.com/kyawzinthant-coding/froky/internal/service/menu"
	"github.com/kyawzinthant-coding/froky/internal/types"
)

const (
	msgUnauthorized      = "unauthorized"
	msgCategoryNotFound  = "category not found"
	msgCategoryNameTaken = "a category with this name already exists"
	msgCategoryNotEmpty  = "category still has products; move or delete them first"
	msgProductNotFound   = "product not found"
	msgProductNameTaken  = "a product with this name already exists"
)

type menuService interface {
	ListCategories(ctx context.Context, vendorID int64) ([]types.Category, error)
	GetCategory(ctx context.Context, vendorID, categoryID int64) (*types.Category, error)
	CreateCategory(ctx context.Context, vendorID int64, name string) (*types.Category, error)
	UpdateCategory(ctx context.Context, vendorID, categoryID int64, name string) (*types.Category, error)
	ArchiveCategory(ctx context.Context, vendorID, categoryID int64) error

	ListProducts(ctx context.Context, vendorID int64) ([]types.Product, error)
	GetProduct(ctx context.Context, vendorID, productID int64) (*types.Product, error)
	CreateProduct(ctx context.Context, vendorID int64, in types.ProductInput) (*types.Product, error)
	UpdateProduct(ctx context.Context, vendorID, productID int64, in types.ProductInput) (*types.Product, error)
	ArchiveProduct(ctx context.Context, vendorID, productID int64) error
}

type VendorMenuHandler struct {
	menuService menuService
}

func NewVendorMenuHandler(svc *menu.Service) *VendorMenuHandler {
	return &VendorMenuHandler{menuService: svc}
}

// The vendor always comes from the session, never from the request body or path.
func vendorIDFromContext(ctx context.Context) (int64, bool) {
	claims, ok := middleware.VendorClaimsFromContext(ctx)
	if !ok {
		return 0, false
	}
	return claims.VendorID, true
}

func unauthorized() api.UnauthorizedJSONResponse {
	return api.UnauthorizedJSONResponse{Error: msgUnauthorized}
}

func toAPICategory(c types.Category) api.Category {
	return api.Category{Id: c.ID, Name: c.Name}
}

func toAPIProduct(p types.Product) api.Product {
	return api.Product{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		ImageUrl:    p.ImageURL,
		CategoryId:  p.CategoryID,
		IsAvailable: p.IsAvailable,
	}
}

func toProductInput(body *api.ProductRequest) types.ProductInput {
	return types.ProductInput{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		ImageURL:    body.ImageUrl,
		CategoryID:  body.CategoryId,
		IsAvailable: body.IsAvailable,
	}
}
