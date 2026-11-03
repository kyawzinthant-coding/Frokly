package database

import (
	"database/sql"
	"github.com/kyawzinthant-coding/froky/config"
	"net"
	"net/url"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func OpenPostgres(cfg *config.DatabaseConfig) (*sql.DB, error) {
	dsn := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.UserName, cfg.Password),
		Host:   net.JoinHostPort(cfg.Host, cfg.Port),
		Path:   cfg.Name,
	}

	query := dsn.Query()
	query.Set("sslmode", "disable")
	dsn.RawQuery = query.Encode()

	return OpenDSN("pgx", dsn.String())
}
