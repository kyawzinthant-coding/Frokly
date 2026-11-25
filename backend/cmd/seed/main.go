package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/kyawzinthant-coding/froky/config"
	"github.com/kyawzinthant-coding/froky/internal/database"
	"golang.org/x/crypto/bcrypt"
)

const foodCourtName = "Central Food Court"

type seedProduct struct {
	name        string
	description string
	price       int64 // satang
	available   bool
}

type seedCategory struct {
	name     string
	products []seedProduct
}

type seedStall struct {
	name     string
	email    string
	username string
	menu     []seedCategory
}

var stalls = []seedStall{
	{
		name: "Khun Mae Noodles", email: "khunmae@forkly.test", username: "vendor1",
		menu: []seedCategory{
			{name: "Noodle Soups", products: []seedProduct{
				{"Tom Yum Noodles", "Thin rice noodles, minced pork, ground peanuts, lime and chilli", 7000, true},
				{"Boat Noodles", "Rich dark pork broth with morning glory and crispy pork rind", 6000, true},
				{"Wonton Egg Noodle Soup", "Egg noodles, shrimp wontons and choy sum in clear broth", 6500, true},
			}},
			{name: "Stir-fried Noodles", products: []seedProduct{
				{"Pad Thai with Shrimp", "Rice noodles, egg, tofu, bean sprouts and peanuts", 8000, true},
				{"Pad See Ew", "Wide rice noodles, Chinese broccoli, dark soy and pork", 7000, true},
				{"Rad Na", "Wide noodles in thick gravy with kale and pork", 7000, false},
			}},
			{name: "Drinks", products: []seedProduct{
				{"Thai Iced Tea", "Sweet, with condensed milk", 3500, true},
				{"Fresh Lime Soda", "", 3000, true},
				{"Bottled Water", "", 1500, true},
			}},
		},
	},
	{
		name: "Baan Khao Gaeng", email: "baankhaogaeng@forkly.test", username: "vendor2",
		menu: []seedCategory{
			{name: "Rice Dishes", products: []seedProduct{
				{"Khao Man Gai", "Poached chicken on garlic rice with ginger-soy sauce and clear soup", 6000, true},
				{"Pad Kra Pao Moo Saap", "Minced pork with holy basil and chilli, fried egg on rice", 6500, true},
				{"Khao Moo Daeng", "Red roast pork and crispy pork belly with sweet gravy", 6000, true},
			}},
			{name: "Curries", products: []seedProduct{
				{"Green Curry Chicken", "With Thai eggplant and basil, served with rice", 7500, true},
				{"Massaman Beef", "Slow-cooked beef, potato and peanuts, served with rice", 9000, false},
				{"Panang Pork", "Thick, mild red curry with kaffir lime, served with rice", 7500, true},
			}},
			{name: "Desserts", products: []seedProduct{
				{"Mango Sticky Rice", "Nam dok mai mango with salted coconut cream", 8000, true},
				{"Khanom Thuay", "Steamed pandan and coconut custard cups", 3500, true},
			}},
		},
	},
	{
		name: "Somtam Corner", email: "somtamcorner@forkly.test", username: "vendor3",
		menu: []seedCategory{
			{name: "Som Tam", products: []seedProduct{
				{"Som Tam Thai", "Green papaya, peanuts, dried shrimp and lime", 5500, true},
				{"Som Tam Pu Pla Ra", "Salted crab and fermented fish sauce, Isaan style", 6000, true},
			}},
			{name: "Grill", products: []seedProduct{
				{"Gai Yang (half chicken)", "Charcoal-grilled, with jaew dipping sauce", 12000, true},
				{"Moo Ping (3 skewers)", "Grilled marinated pork skewers", 4500, true},
				{"Kor Moo Yang", "Grilled pork neck with jaew", 9000, true},
			}},
			{name: "Sides", products: []seedProduct{
				{"Larb Moo", "Minced pork, toasted rice, mint and lime", 7000, true},
				{"Sticky Rice", "", 1500, true},
			}},
		},
	},
}

func main() {
	password := flag.String("password", "password123", "password for staff accounts this run creates (plaintext, will be hashed)")
	flag.Parse()

	cfg := config.LoadConfig()

	db, err := database.OpenPostgres(cfg.DataBase)
	if err != nil {
		log.Fatal("open db: ", err)
	}
	defer db.Close()

	if err := run(db, *password); err != nil {
		log.Fatal("seed failed: ", err)
	}

	fmt.Println("seed completed")
}

// run is safe to repeat: rows that already exist are reused or skipped, never overwritten,
// so menu edits made in the app survive a re-seed.
func run(db *sql.DB, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	foodCourtID, err := ensureFoodCourt(tx)
	if err != nil {
		return err
	}

	for _, stall := range stalls {
		vendorID, err := ensureStall(tx, foodCourtID, stall, string(hash))
		if err != nil {
			return fmt.Errorf("stall %q: %w", stall.name, err)
		}
		if err := seedMenu(tx, vendorID, stall.menu); err != nil {
			return fmt.Errorf("menu of %q: %w", stall.name, err)
		}
	}

	return tx.Commit()
}

func ensureFoodCourt(tx *sql.Tx) (int64, error) {
	var id int64
	err := tx.QueryRow(`SELECT id FROM food_courts WHERE name = $1 ORDER BY id LIMIT 1`, foodCourtName).Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		err = tx.QueryRow(`INSERT INTO food_courts (name) VALUES ($1) RETURNING id`, foodCourtName).Scan(&id)
	}
	if err != nil {
		return 0, fmt.Errorf("food court: %w", err)
	}
	return id, nil
}

// The staff username is the key: an existing login keeps its password and gets the stall's
// name and email; a new username gets a new vendor and staff account.
func ensureStall(tx *sql.Tx, foodCourtID int64, stall seedStall, passwordHash string) (int64, error) {
	var vendorID int64
	err := tx.QueryRow(`SELECT vendor_id FROM vendor_staff WHERE username = $1`, stall.username).Scan(&vendorID)
	if err == nil {
		_, err = tx.Exec(
			`UPDATE vendors SET name = $2, email = $3, updated_at = NOW() WHERE id = $1`,
			vendorID, stall.name, stall.email,
		)
		return vendorID, err
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}

	err = tx.QueryRow(
		`INSERT INTO vendors (name, email, food_court_id) VALUES ($1, $2, $3) RETURNING id`,
		stall.name, stall.email, foodCourtID,
	).Scan(&vendorID)
	if err != nil {
		return 0, err
	}

	_, err = tx.Exec(
		`INSERT INTO vendor_staff (vendor_id, username, password_hash) VALUES ($1, $2, $3)`,
		vendorID, stall.username, passwordHash,
	)
	return vendorID, err
}

func seedMenu(tx *sql.Tx, vendorID int64, menu []seedCategory) error {
	for _, category := range menu {
		var categoryID int64
		// The no-op DO UPDATE makes RETURNING give back the existing row's id on a conflict.
		err := tx.QueryRow(`
			INSERT INTO menu_categories (vendor_id, name) VALUES ($1, $2)
			ON CONFLICT (vendor_id, lower(name)) WHERE archived_at IS NULL
			DO UPDATE SET name = menu_categories.name
			RETURNING id`,
			vendorID, category.name,
		).Scan(&categoryID)
		if err != nil {
			return fmt.Errorf("category %q: %w", category.name, err)
		}

		for _, p := range category.products {
			_, err := tx.Exec(`
				INSERT INTO products (name, description, price, menu_categories_id, vendor_id, is_available)
				VALUES ($1, $2, $3, $4, $5, $6)
				ON CONFLICT (vendor_id, lower(name)) WHERE archived_at IS NULL DO NOTHING`,
				p.name, p.description, p.price, categoryID, vendorID, p.available,
			)
			if err != nil {
				return fmt.Errorf("product %q: %w", p.name, err)
			}
		}
	}
	return nil
}
