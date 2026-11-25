package menu

import (
	"context"
	"database/sql"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	dbgen "github.com/kyawzinthant-coding/froky/db/generated"
	"github.com/kyawzinthant-coding/froky/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testDatabaseURLEnv = "FORKLY_TEST_DATABASE_URL"

// newTxService runs the service inside a transaction that is rolled back after the
// test, so nothing is left in the database. Skipped unless FORKLY_TEST_DATABASE_URL is set.
func newTxService(t *testing.T) (*Service, *sql.Tx) {
	t.Helper()

	dsn := os.Getenv(testDatabaseURLEnv)
	if dsn == "" {
		t.Skipf("%s not set", testDatabaseURLEnv)
	}

	db, err := sql.Open("pgx", dsn)
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })

	tx, err := db.Begin()
	require.NoError(t, err)
	t.Cleanup(func() { _ = tx.Rollback() })

	return NewService(dbgen.New(tx)), tx
}

func insertVendor(t *testing.T, tx *sql.Tx, name string) int64 {
	t.Helper()

	var foodCourtID, vendorID int64
	require.NoError(t, tx.QueryRow(`INSERT INTO food_courts (name) VALUES ('Test Court') RETURNING id`).Scan(&foodCourtID))
	require.NoError(t, tx.QueryRow(
		`INSERT INTO vendors (name, email, food_court_id) VALUES ($1, $1 || '@test.local', $2) RETURNING id`,
		name, foodCourtID,
	).Scan(&vendorID))
	return vendorID
}

func TestMenu_VendorIsolation(t *testing.T) {
	svc, tx := newTxService(t)
	ctx := context.Background()

	vendorA := insertVendor(t, tx, "isolation-vendor-a")
	vendorB := insertVendor(t, tx, "isolation-vendor-b")

	category, err := svc.CreateCategory(ctx, vendorA, "Noodles")
	require.NoError(t, err)
	product, err := svc.CreateProduct(ctx, vendorA, types.ProductInput{
		Name: "Pad Thai", Price: 6500, CategoryID: category.ID, IsAvailable: true,
	})
	require.NoError(t, err)

	t.Run("B cannot read A's menu", func(t *testing.T) {
		_, err := svc.GetProduct(ctx, vendorB, product.ID)
		assert.ErrorIs(t, err, ErrProductNotFound)

		_, err = svc.GetCategory(ctx, vendorB, category.ID)
		assert.ErrorIs(t, err, ErrCategoryNotFound)

		products, err := svc.ListProducts(ctx, vendorB)
		require.NoError(t, err)
		assert.Empty(t, products)
	})

	t.Run("B cannot change A's menu", func(t *testing.T) {
		_, err := svc.UpdateProduct(ctx, vendorB, product.ID, types.ProductInput{Name: "Hacked", Price: 1, CategoryID: category.ID})
		assert.ErrorIs(t, err, ErrProductNotFound)

		_, err = svc.UpdateCategory(ctx, vendorB, category.ID, "Hacked")
		assert.ErrorIs(t, err, ErrCategoryNotFound)

		assert.ErrorIs(t, svc.ArchiveProduct(ctx, vendorB, product.ID), ErrProductNotFound)
		assert.ErrorIs(t, svc.ArchiveCategory(ctx, vendorB, category.ID), ErrCategoryNotFound)
	})

	t.Run("B cannot put a product in A's category", func(t *testing.T) {
		_, err := svc.CreateProduct(ctx, vendorB, types.ProductInput{Name: "Sneaky", Price: 1, CategoryID: category.ID})
		assert.ErrorIs(t, err, ErrCategoryNotFound)
	})

	t.Run("A's menu is unchanged", func(t *testing.T) {
		got, err := svc.GetProduct(ctx, vendorA, product.ID)
		require.NoError(t, err)
		assert.Equal(t, "Pad Thai", got.Name)
	})
}

func TestMenu_ArchiveRules(t *testing.T) {
	svc, tx := newTxService(t)
	ctx := context.Background()
	vendor := insertVendor(t, tx, "archive-vendor")

	category, err := svc.CreateCategory(ctx, vendor, "Noodles")
	require.NoError(t, err)
	product, err := svc.CreateProduct(ctx, vendor, types.ProductInput{Name: "Pad Thai", Price: 6500, CategoryID: category.ID})
	require.NoError(t, err)

	assert.ErrorIs(t, svc.ArchiveCategory(ctx, vendor, category.ID), ErrCategoryNotEmpty)

	require.NoError(t, svc.ArchiveProduct(ctx, vendor, product.ID))
	_, err = svc.GetProduct(ctx, vendor, product.ID)
	assert.ErrorIs(t, err, ErrProductNotFound)

	_, err = svc.CreateProduct(ctx, vendor, types.ProductInput{Name: "pad thai", Price: 7000, CategoryID: category.ID})
	require.NoError(t, err, "an archived name can be reused")

	// Runs last: a unique violation aborts the surrounding transaction.
	_, err = svc.CreateProduct(ctx, vendor, types.ProductInput{Name: "PAD THAI", Price: 1, CategoryID: category.ID})
	assert.ErrorIs(t, err, ErrNameTaken, "names are unique case-insensitively")
}
