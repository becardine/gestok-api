-- name: CreateProductStock :exec
INSERT INTO product_stocks (
    id, product_id, stock_id, quantity, created_at, updated_at
)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetProductStock :one
SELECT *
FROM product_stocks
WHERE product_id = ? AND stock_id = ? AND deleted_at IS NULL;

-- name: UpdateProductStock :exec
UPDATE product_stocks
SET quantity = ?, updated_at = ?
WHERE product_id = ? AND stock_id = ? AND deleted_at IS NULL;

-- name: DeleteProductStock :exec
UPDATE product_stocks
SET deleted_at = NOW()
WHERE product_id = ? AND stock_id = ?;

-- name: ListProductStocks :many
SELECT *
FROM product_stocks
WHERE product_id = ? AND deleted_at IS NULL
ORDER BY created_at
LIMIT 2 OFFSET 3;

-- name: ListProductsInStock :many
SELECT p.*
FROM products p
JOIN product_stocks ps ON p.id = ps.product_id
WHERE ps.stock_id = ? AND ps.deleted_at IS NULL AND p.deleted_at IS NULL
ORDER BY p.name
LIMIT 2 OFFSET 3;
