package auth

import (
	"context"
	"database/sql"
	"errors"
	"github.com/kyawzinthant-coding/froky/internal/types"
)

var ErrStaffNotFound = errors.New("staff not found")

func (s *Service) Me(ctx context.Context, staffId int64) (*types.VendorProfile, error) {
	row, err := s.q.GetVendorProfileByStaffId(ctx, staffId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrStaffNotFound
	}
	if err != nil {
		return nil, err
	}

	return &types.VendorProfile{
		ID:       row.VendorID,
		Name:     row.VendorName,
		Username: row.Username,
	}, nil
}
