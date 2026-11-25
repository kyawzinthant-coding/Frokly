package menu

import (
	"context"

	"github.com/kyawzinthant-coding/froky/internal/types"
)

// AvailableMenu is what a customer sees: active categories that have at least one item
// on sale, each with only its on-sale items. Two queries, grouped here, so the number
// of queries doesn't grow with the number of categories.
func (s *Service) AvailableMenu(ctx context.Context, vendorID int64) ([]types.MenuSection, error) {
	categories, err := s.q.ListCategoriesByVendor(ctx, vendorID)
	if err != nil {
		return nil, err
	}

	products, err := s.q.ListAvailableProductsByVendor(ctx, vendorID)
	if err != nil {
		return nil, err
	}

	itemsByCategory := make(map[int64][]types.Product, len(categories))
	for _, row := range products {
		itemsByCategory[row.MenuCategoriesID] = append(itemsByCategory[row.MenuCategoriesID], toProduct(row))
	}

	// Products in an archived category are dropped here too: only active categories are walked.
	sections := make([]types.MenuSection, 0, len(categories))
	for _, category := range categories {
		items := itemsByCategory[category.ID]
		if len(items) == 0 {
			continue
		}
		sections = append(sections, types.MenuSection{ID: category.ID, Name: category.Name, Items: items})
	}
	return sections, nil
}
