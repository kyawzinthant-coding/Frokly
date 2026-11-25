package menu

import (
	"context"
	"database/sql"
	"errors"

	dbgen "github.com/kyawzinthant-coding/froky/db/generated"
	"github.com/kyawzinthant-coding/froky/internal/types"
)

func (s *Service) ListCategories(ctx context.Context, vendorID int64) ([]types.Category, error) {
	rows, err := s.q.ListCategoriesByVendor(ctx, vendorID)
	if err != nil {
		return nil, err
	}

	categories := make([]types.Category, 0, len(rows))
	for _, row := range rows {
		categories = append(categories, toCategory(row))
	}
	return categories, nil
}

func (s *Service) GetCategory(ctx context.Context, vendorID, categoryID int64) (*types.Category, error) {
	row, err := s.q.GetCategoryByVendor(ctx, dbgen.GetCategoryByVendorParams{ID: categoryID, VendorID: vendorID})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrCategoryNotFound
	}
	if err != nil {
		return nil, err
	}

	category := toCategory(row)
	return &category, nil
}

func (s *Service) CreateCategory(ctx context.Context, vendorID int64, name string) (*types.Category, error) {
	name, err := normalizeName(name)
	if err != nil {
		return nil, err
	}

	row, err := s.q.CreateCategory(ctx, dbgen.CreateCategoryParams{VendorID: vendorID, Name: name})
	if isUniqueViolation(err) {
		return nil, ErrNameTaken
	}
	if err != nil {
		return nil, err
	}

	category := toCategory(row)
	return &category, nil
}

func (s *Service) UpdateCategory(ctx context.Context, vendorID, categoryID int64, name string) (*types.Category, error) {
	name, err := normalizeName(name)
	if err != nil {
		return nil, err
	}

	row, err := s.q.UpdateCategory(ctx, dbgen.UpdateCategoryParams{ID: categoryID, VendorID: vendorID, Name: name})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrCategoryNotFound
	}
	if isUniqueViolation(err) {
		return nil, ErrNameTaken
	}
	if err != nil {
		return nil, err
	}

	category := toCategory(row)
	return &category, nil
}

func (s *Service) ArchiveCategory(ctx context.Context, vendorID, categoryID int64) error {
	archived, err := s.q.ArchiveEmptyCategory(ctx, dbgen.ArchiveEmptyCategoryParams{ID: categoryID, VendorID: vendorID})
	if err != nil {
		return err
	}
	if archived > 0 {
		return nil
	}

	// Nothing archived: either the category isn't this vendor's, or it still has products.
	if _, err := s.GetCategory(ctx, vendorID, categoryID); err != nil {
		return err
	}
	return ErrCategoryNotEmpty
}

func toCategory(row dbgen.MenuCategory) types.Category {
	return types.Category{ID: row.ID, Name: row.Name}
}
