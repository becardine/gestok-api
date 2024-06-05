-- name: GetBrand :one
SELECT * FROM brands WHERE id = $1 AND deleted_at IS NULL;

-- name: CreateBrand :exec
INSERT INTO brands (id, name, description, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateBrand :exec
UPDATE brands
SET name = $2, description = $3, updated_date = $4
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteBrand :exec
UPDATE brands
SET deleted_at = NOW()
WHERE id = $1;

-- name: ListBrands :many
SELECT * FROM brands WHERE deleted_at IS NULL ORDER BY name;

-- name: GetBrandProducts :many
SELECT p.*
FROM brands b
JOIN products p ON b.id = p.brand_id
WHERE b.id = $1 AND b.deleted_at IS NULL
ORDER BY p.name;