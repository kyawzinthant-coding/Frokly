package database

import (
	"context"
	"database/sql"
	"os"

	"github.com/pressly/goose/v3"
)

type Migrator struct {
	provider *goose.Provider
}

func NewMigrator(db *sql.DB, migrationDir string) (*Migrator, error) {
	provider, err := goose.NewProvider(
		goose.Dialect("postgres"),
		db,
		os.DirFS(migrationDir),
	)
	if err != nil {
		return nil, err
	}

	return &Migrator{provider: provider}, nil
}

func (m *Migrator) Up(ctx context.Context) error {
	_, err := m.provider.Up(ctx)
	return err
}

func (m *Migrator) Down(ctx context.Context) error {
	_, err := m.provider.Down(ctx)
	return err
}

func (m *Migrator) Reset(ctx context.Context) error {
	_, err := m.provider.DownTo(ctx, 0)
	return err
}

func (m *Migrator) Status(ctx context.Context) ([]*goose.MigrationStatus, error) {
	return m.provider.Status(ctx)
}
