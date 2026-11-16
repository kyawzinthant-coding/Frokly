-- name: GetVendorStaffForLogin :one
SELECT
    vs.id            AS staff_id,
    vs.username,
    vs.password_hash,
    v.id             AS vendor_id,
    v.name           AS vendor_name
FROM vendor_staff vs
JOIN vendors v ON v.id = vs.vendor_id
WHERE vs.username = @username;
