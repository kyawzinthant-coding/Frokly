package menu

import (
	"context"
	"database/sql"
	"errors"

	dbgen "github.com/kyawzinthant-coding/froky/db/generated"
	"github.com/kyawzinthant-coding/froky/internal/types"
)

func (s *Service) ListProducts(ctx context.Context, vendorID int64) ([]types.Product, error) {
	rows, err := s.q.ListProductsByVendor(ctx, vendorID)
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0, len(rows))
	for _, row := range rows {
		products = append(products, toProduct(row))
	}
	return products, nil
}

func (s *Service) GetProduct(ctx context.Context, vendorID, productID int64) (*types.Product, error) {
	row, err := s.q.GetProductByVendor(ctx, dbgen.GetProductByVendorParams{ID: productID, VendorID: vendorID})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrProductNotFound
	}
	if err != nil {
		return nil, err
	}

	product := toProduct(row)
	return &product, nil
}

func (s *Service) CreateProduct(ctx context.Context, vendorID int64, in types.ProductInput) (*types.Product, error) {
	in, err := normalizeProductInput(in)
	if err != nil {
		return nil, err
	}

	row, err := s.q.CreateProduct(ctx, dbgen.CreateProductParams{
		Name:        in.Name,
		Description: in.Description,
		ProductUrl:  toNullString(in.ImageURL),
		Price:       in.Price,
		IsAvailable: in.IsAvailable,
		CategoryID:  in.CategoryID,
		VendorID:    vendorID,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrCategoryNotFound
	}
	if isUniqueViolation(err) {
		return nil, ErrNameTaken
	}
	if err != nil {
		return nil, err
	}

	product := toProduct(row)
	return &product, nil
}

func (s *Service) UpdateProduct(ctx context.Context, vendorID, productID int64, in types.ProductInput) (*types.Product, error) {
	in, err := normalizeProductInput(in)
	if err != nil {
		return nil, err
	}

	row, err := s.q.UpdateProduct(ctx, dbgen.UpdateProductParams{
		Name:        in.Name,
		Description: in.Description,
		ProductUrl:  toNullString(in.ImageURL),
		Price:       in.Price,
		CategoryID:  in.CategoryID,
		IsAvailable: in.IsAvailable,
		ID:          productID,
		VendorID:    vendorID,
	})
	if errors.Is(err, sql.ErrNoRows) {
		// Nothing updated: either the product or the target category isn't this vendor's.
		if _, err := s.GetProduct(ctx, vendorID, productID); err != nil {
			return nil, err
		}
		return nil, ErrCategoryNotFound
	}
	if isUniqueViolation(err) {
		return nil, ErrNameTaken
	}
	if err != nil {
		return nil, err
	}

	product := toProduct(row)
	return &product, nil
}

func (s *Service) ArchiveProduct(ctx context.Context, vendorID, productID int64) error {
	archived, err := s.q.ArchiveProduct(ctx, dbgen.ArchiveProductParams{ID: productID, VendorID: vendorID})
	if err != nil {
		return err
	}
	if archived == 0 {
		return ErrProductNotFound
	}
	return nil
}

func toProduct(row dbgen.Product) types.Product {
	var imageURL *string
	if row.ProductUrl.Valid {
		imageURL = &row.ProductUrl.String
	}

	return types.Product{
		ID:          row.ID,
		Name:        row.Name,
		Description: row.Description,
		Price:       row.Price,
		ImageURL:    imageURL,
		CategoryID:  row.MenuCategoriesID,
		IsAvailable: row.IsAvailable,
	}
}

func toNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *s, Valid: true}
}
