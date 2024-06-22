-- name: GetBrand :one
SELECT * FROM brands WHERE id = ? AND deleted_at IS NULL;

-- name: CreateBrand :exec
INSERT INTO brands (id, name, description, created_at, updated_at)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateBrand :exec
UPDATE brands
SET name = ?, description = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: DeleteBrand :exec
UPDATE brands
SET deleted_at = NOW()
WHERE id = ?;

-- name: ListBrands :many
SELECT *
FROM brands
WHERE deleted_at IS NULL
ORDER BY name
LIMIT ? OFFSET ?;

-- name: GetBrandProducts :many
SELECT p.*
FROM brands b
JOIN products p ON b.id = p.brand_id
WHERE b.id = ? AND b.deleted_at IS NULL AND p.deleted_at IS NULL
ORDER BY p.name;