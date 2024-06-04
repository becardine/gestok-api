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
SELECT * FROM stocks WHERE deleted_at IS NULL ORDER BY name;

-- name: GetStockProducts :many
SELECT p.*
FROM stocks s
JOIN product_stocks ps ON s.id = ps.stock_id
JOIN products p ON ps.product_id = p.id
WHERE s.id = $1 AND s.deleted_at IS NULL;

-- name: AddStockProduct :exec
INSERT INTO product_stocks (stock_id, product_id)
VALUES ($1, $2);

-- name: RemoveStockProduct :exec
DELETE FROM product_stocks WHERE stock_id = $1 AND product_id = $2;