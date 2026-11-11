package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/kyawzinthant-coding/froky/config"
	"github.com/kyawzinthant-coding/froky/internal/database"
	"log"
	"time"
)

func main() {
	action := flag.String("action", "up", "action to perform migration")

	flag.Parse()

	if err := run(*action); err != nil {
		log.Fatal("migration failed", err)
	}
}

func run(action string) error {

	cfg := config.LoadConfig()

	db, err := database.OpenPostgres(cfg.DataBase)
	if err != nil {
		return err
	}
	defer db.Close()

	migrator, err := database.NewMigrator(db, "db/migrations/version1")
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	switch action {
	case "up":
		if err := migrator.Up(ctx); err != nil {
			return fmt.Errorf("migrate up: %w", err)
		}
		fmt.Println("migrations up completed")

	case "down":
		if err := migrator.Down(ctx); err != nil {
			return fmt.Errorf("migrate down: %w", err)
		}
		fmt.Println("one migration rolled back")

	case "reset":
		if err := migrator.Reset(ctx); err != nil {
			return fmt.Errorf("migrate reset: %w", err)
		}
		fmt.Println("all migrations rolled back")

	case "status":
		statuses, err := migrator.Status(ctx)
		if err != nil {
			return fmt.Errorf("migration status: %w", err)
		}

		for _, status := range statuses {
			fmt.Printf(
				"%-10v version=%d file=%s\n",
				status.State,
				status.Source.Version,
				status.Source.Path,
			)
		}

	default:
		return fmt.Errorf("unknown action %q; use up, down, reset, or status", action)
	}

	return nil

}
