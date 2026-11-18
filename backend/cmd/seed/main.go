package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/kyawzinthant-coding/froky/config"
	"github.com/kyawzinthant-coding/froky/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	vendorName := flag.String("vendor-name", "Test Vendor", "vendor business name")
	vendorEmail := flag.String("vendor-email", "vendor@example.com", "vendor contact email")
	username := flag.String("username", "vendor1", "vendor_staff login username")
	password := flag.String("password", "password123", "vendor_staff login password (plaintext, will be hashed)")

	flag.Parse()

	cfg := config.LoadConfig()

	db, err := database.OpenPostgres(cfg.DataBase)
	if err != nil {
		log.Fatal("open db: ", err)
	}
	defer db.Close()

	if err := run(db, *vendorName, *vendorEmail, *username, *password); err != nil {
		log.Fatal("seed failed: ", err)
	}

	fmt.Println("seed completed")
}

func run(db *sql.DB, vendorName, vendorEmail, username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	var foodCourtID int64
	err = tx.QueryRow(`
		INSERT INTO food_courts (name)
		VALUES ('Central Food Court')
		RETURNING id
	`).Scan(&foodCourtID)
	if err != nil {
		return fmt.Errorf("insert food court: %w", err)
	}

	var vendorID int64
	err = tx.QueryRow(`
		INSERT INTO vendors (name, email, food_court_id)
		VALUES ($1, $2, $3)
		RETURNING id
	`, vendorName, vendorEmail, foodCourtID).Scan(&vendorID)
	if err != nil {
		return fmt.Errorf("insert vendor: %w", err)
	}

	_, err = tx.Exec(`
		INSERT INTO vendor_staff (vendor_id, username, password_hash)
		VALUES ($1, $2, $3)
	`, vendorID, username, string(hash))
	if err != nil {
		return fmt.Errorf("insert vendor_staff: %w", err)
	}

	return tx.Commit()
}
