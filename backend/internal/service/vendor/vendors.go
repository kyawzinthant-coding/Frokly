package vendor

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kyawzinthant-coding/froky/internal/nullable"
	"github.com/kyawzinthant-coding/froky/internal/types"
)

// List returns every stall. There is one food court in the MVP; once table sessions
// exist this must be limited to the food court of the customer's table.
func (s *Service) List(ctx context.Context) ([]types.Vendor, error) {
	rows, err := s.q.ListVendors(ctx)
	if err != nil {
		return nil, err
	}

	vendors := make([]types.Vendor, 0, len(rows))
	for _, row := range rows {
		vendors = append(vendors, types.Vendor{ID: row.ID, Name: row.Name, LogoURL: nullable.StringPtr(row.Logourl)})
	}
	return vendors, nil
}

func (s *Service) Get(ctx context.Context, vendorID int64) (*types.Vendor, error) {
	row, err := s.q.GetVendor(ctx, vendorID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrVendorNotFound
	}
	if err != nil {
		return nil, err
	}

	return &types.Vendor{ID: row.ID, Name: row.Name, LogoURL: nullable.StringPtr(row.Logourl)}, nil
}
