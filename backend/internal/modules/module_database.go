package modules

import (
	"context"
	"database/sql"

	dbgen "github.com/kyawzinthant-coding/froky/db/generated"
	"github.com/kyawzinthant-coding/froky/internal/database"
	"go.uber.org/fx"
)

var Database = fx.Module("database",
	fx.Provide(database.OpenPostgres),
	// dbgen.New takes the DBTX interface, but fx matches on exact types,
	// so this small adapter turns the *sql.DB fx already has into *dbgen.Queries.
	fx.Provide(func(db *sql.DB) *dbgen.Queries { return dbgen.New(db) }),
	fx.Invoke(func(lc fx.Lifecycle, db *sql.DB) {
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				return db.Close()
			},
		})
	}),
)
