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

	var staffID int64
	var passwordHash string
	var profile types.VendorProfile

	err := s.store.QueryRowContext(ctx, `
                SELECT vs.id, vs.password_hash, v.id, v.name, vs.username
                FROM vendor_staff vs
                JOIN vendors v ON v.id = vs.vendor_id
                WHERE vs.username = $1`,
		username,
	).Scan(&staffID, &passwordHash, &profile.ID, &profile.Name, &profile.Username)

	if errors.Is(err, sql.ErrNoRows) {
		return "", nil, ErrInvalidCredentials
	}
	if err != nil {
		return "", nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) != nil {
		return "", nil, ErrInvalidCredentials
	}

	now := time.Now()
	claims := types.VendorClaims{
		VendorID: profile.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatInt(staffID, 10),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.expiry)),
		},
	}

	signed, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.jwtSecret)
	if err != nil {
		return "", nil, err
	}

	return signed, &profile, nil
}
