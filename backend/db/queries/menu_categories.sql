-- name: ListCategoriesByVendor :many
SELECT * FROM menu_categories
WHERE vendor_id = @vendor_id AND archived_at IS NULL
ORDER BY name;

-- name: GetCategoryByVendor :one
SELECT * FROM menu_categories
WHERE id = @id AND vendor_id = @vendor_id AND archived_at IS NULL;

-- name: CreateCategory :one
INSERT INTO menu_categories (vendor_id, name)
VALUES (@vendor_id, @name)
RETURNING *;

-- name: UpdateCategory :one
UPDATE menu_categories
SET name = @name, updated_at = NOW()
WHERE id = @id AND vendor_id = @vendor_id AND archived_at IS NULL
RETURNING *;

-- name: ArchiveEmptyCategory :execrows
UPDATE menu_categories c
SET archived_at = NOW(), updated_at = NOW()
WHERE c.id = @id AND c.vendor_id = @vendor_id AND c.archived_at IS NULL
  AND NOT EXISTS (
      SELECT 1 FROM products p
      WHERE p.menu_categories_id = c.id AND p.archived_at IS NULL
  );
