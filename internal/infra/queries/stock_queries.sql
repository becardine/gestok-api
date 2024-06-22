-- name: GetStock :one
SELECT * FROM stocks WHERE id = ? AND deleted_at IS NULL;

-- name: CreateStock :exec
INSERT INTO stocks (id, name, location, capacity, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateStock :exec
UPDATE stocks
SET name = ?, location = ?, capacity = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: DeleteStock :exec
UPDATE stocks
SET deleted_at = NOW()
WHERE id = ?;

-- name: ListStocks :many
SELECT *
FROM stocks
WHERE deleted_at IS NULL
ORDER BY name
LIMIT 1 OFFSET 2;

-- name: GetStockProducts :many
SELECT p.*
FROM stocks s
JOIN product_stocks ps ON s.id = ps.stock_id
JOIN products p ON ps.product_id = p.id
WHERE s.id = ? AND s.deleted_at IS NULL AND p.deleted_at IS NULL
ORDER BY p.name;

-- name: AddStockProduct :exec
INSERT INTO product_stocks (stock_id, product_id)
VALUES (?, ?)
ON DUPLICATE KEY UPDATE stock_id = stock_id;

-- name: RemoveStockProduct :execresult
DELETE FROM product_stocks
WHERE stock_id = ? AND product_id = ?;
