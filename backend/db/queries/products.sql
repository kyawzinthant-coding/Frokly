-- name: ListProductsByVendor :many
SELECT * FROM products
WHERE vendor_id = @vendor_id AND archived_at IS NULL
ORDER BY name;

-- name: GetProductByVendor :one
SELECT * FROM products
WHERE id = @id AND vendor_id = @vendor_id AND archived_at IS NULL;

-- name: CreateProduct :one
INSERT INTO products (name, description, product_url, price, menu_categories_id, vendor_id, is_available)
SELECT @name::text, @description::text, sqlc.narg(product_url)::text, @price::bigint,
       c.id, c.vendor_id, @is_available::boolean
FROM menu_categories c
WHERE c.id = @category_id AND c.vendor_id = @vendor_id AND c.archived_at IS NULL
RETURNING *;

-- name: UpdateProduct :one
UPDATE products
SET name               = @name,
    description        = @description,
    product_url        = sqlc.narg(product_url),
    price              = @price,
    menu_categories_id = @category_id,
    is_available       = @is_available,
    updated_at         = NOW()
WHERE products.id = @id AND products.vendor_id = @vendor_id AND products.archived_at IS NULL
  AND EXISTS (
      SELECT 1 FROM menu_categories c
      WHERE c.id = @category_id AND c.vendor_id = @vendor_id AND c.archived_at IS NULL
  )
RETURNING *;

-- name: ArchiveProduct :execrows
UPDATE products
SET archived_at = NOW(), updated_at = NOW()
WHERE id = @id AND vendor_id = @vendor_id AND archived_at IS NULL;

-- name: ListAvailableProductsByVendor :many
SELECT * FROM products
WHERE vendor_id = @vendor_id AND is_available AND archived_at IS NULL
ORDER BY name;
