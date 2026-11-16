package auth

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kyawzinthant-coding/froky/internal/types"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

func (s *Service) Login(ctx context.Context, username, password string) (string, *types.VendorProfile, error) {
	row, err := s.q.GetVendorStaffForLogin(ctx, username)
	if errors.Is(err, sql.ErrNoRows) {
		return "", nil, ErrInvalidCredentials
	}
	if err != nil {
		return "", nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(row.PasswordHash), []byte(password)) != nil {
		return "", nil, ErrInvalidCredentials
	}

	now := time.Now()
	claims := types.VendorClaims{
		VendorID: row.VendorID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatInt(row.StaffID, 10),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.expiry)),
		},
	}

	signed, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.jwtSecret)
	if err != nil {
		return "", nil, err
	}

	return signed, &types.VendorProfile{
		ID:       row.VendorID,
		Name:     row.VendorName,
		Username: row.Username,
	}, nil
}
