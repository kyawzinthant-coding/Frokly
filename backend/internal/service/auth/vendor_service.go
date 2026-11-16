package auth

import (
	"database/sql"
	"github.com/kyawzinthant-coding/froky/config"
	"time"
)

type Service struct {
	store     *sql.DB
	jwtSecret []byte
	expiry    time.Duration
}

func NewService(cfg *config.AppConfig, store *sql.DB) *Service {
	return &Service{
		store:     store,
		jwtSecret: []byte(cfg.JWT.Secret),
		expiry:    cfg.JWT.Expiry,
	}
}
