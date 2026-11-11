package modules

import (
	"context"
	"database/sql"
	"github.com/kyawzinthant-coding/froky/internal/database"
	"go.uber.org/fx"
)

var Database = fx.Module("database",
	fx.Provide(database.OpenPostgres),
	fx.Invoke(func(lc fx.Lifecycle, db *sql.DB) {
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				return db.Close()
			},
		})
	}),
)
