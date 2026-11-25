package public_menu

import (
	"context"

	"github.com/kyawzinthant-coding/froky/api"
	"github.com/kyawzinthant-coding/froky/internal/service/menu"
	"github.com/kyawzinthant-coding/froky/internal/service/vendor"
	"github.com/kyawzinthant-coding/froky/internal/types"
)

const msgVendorNotFound = "vendor not found"

type vendorService interface {
	List(ctx context.Context) ([]types.Vendor, error)
	Get(ctx context.Context, vendorID int64) (*types.Vendor, error)
}

type menuService interface {
	AvailableMenu(ctx context.Context, vendorID int64) ([]types.MenuSection, error)
}

// PublicMenuHandler serves customers, who have no account: no session is read here.
type PublicMenuHandler struct {
	vendorService vendorService
	menuService   menuService
}

func NewPublicMenuHandler(vendors *vendor.Service, menus *menu.Service) *PublicMenuHandler {
	return &PublicMenuHandler{vendorService: vendors, menuService: menus}
}

func toPublicVendor(v types.Vendor) api.PublicVendor {
	return api.PublicVendor{Id: v.ID, Name: v.Name, LogoUrl: v.LogoURL}
}

func toPublicCategory(section types.MenuSection) api.PublicMenuCategory {
	items := make([]api.PublicMenuItem, 0, len(section.Items))
	for _, p := range section.Items {
		items = append(items, api.PublicMenuItem{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			ImageUrl:    p.ImageURL,
		})
	}
	return api.PublicMenuCategory{Id: section.ID, Name: section.Name, Items: items}
}
