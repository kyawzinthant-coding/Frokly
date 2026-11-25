-- name: ListVendors :many
SELECT id, name, logourl FROM vendors
ORDER BY name;

-- name: GetVendor :one
SELECT id, name, logourl FROM vendors
WHERE id = @id;
