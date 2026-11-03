package database

import (
	"database/sql"
	"fmt"
)

func OpenDSN(driverName, dsn string) (*sql.DB, error) {
	database, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	if err := database.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	return database, nil
}
