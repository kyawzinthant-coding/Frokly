package types

import "github.com/golang-jwt/jwt/v5"

type VendorClaims struct {
	VendorID int64 `json:"vendor_id"`
	jwt.RegisteredClaims
}

type VendorProfile struct {
	ID       int64
	Name     string
	Username string
}
