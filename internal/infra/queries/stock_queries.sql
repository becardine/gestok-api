-- name: GetStock :one
SELECT * FROM stocks WHERE id = $1 AND deleted_at IS NULL;

-- name: CreateStock :exec
INSERT INTO stocks (id, name, location, capacity, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpdateStock :exec
UPDATE stocks
SET name = $2, location = $3, capacity = $4, updated_date = $5
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteStock :exec
UPDATE stocks
SET deleted_at = NOW()
WHERE id = $1;

-- name: ListStocks :many
SELECT *
FROM stocks
WHERE deleted_at IS NULL
ORDER BY name
LIMIT $1 OFFSET $2; -- Paginação

-- name: GetStockProducts :many
SELECT p.*
FROM stocks s
         JOIN product_stocks ps ON s.id = ps.stock_id
         JOIN products p ON ps.product_id = p.id
WHERE s.id = $1 AND s.deleted_at IS NULL AND p.deleted_at IS NULL
ORDER BY p.name;

-- name: AddStockProduct :exec
INSERT INTO product_stocks (stock_id, product_id)
VALUES ($1, $2)
ON CONFLICT (stock_id, product_id) DO NOTHING;

-- name: RemoveStockProduct :execresult
DELETE FROM product_stocks
WHERE stock_id = $1 AND product_id = $2
RETURNING *;