package auth

import (
	"time"

	"github.com/kyawzinthant-coding/froky/config"
	dbgen "github.com/kyawzinthant-coding/froky/db/generated"
)

type Service struct {
	q         *dbgen.Queries
	jwtSecret []byte
	expiry    time.Duration
}

func NewService(cfg *config.AppConfig, q *dbgen.Queries) *Service {
	return &Service{
		q:         q,
		jwtSecret: []byte(cfg.JWT.Secret),
		expiry:    cfg.JWT.Expiry,
	}
}
