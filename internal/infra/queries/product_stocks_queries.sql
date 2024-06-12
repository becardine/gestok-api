-- name: CreateProductStock :one
INSERT INTO product_stock (
  id, product_id, stock_id, quantity, created_at, updated_at
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetProductStock :one
SELECT * 
FROM product_stock 
WHERE product_id = $1 AND stock_id = $2 AND deleted_at IS NULL;

-- name: UpdateProductStock :exec
UPDATE product_stock
SET quantity = $3, updated_at = $4
WHERE product_id = $1 AND stock_id = $2 AND deleted_at IS NULL;

-- name: DeleteProductStock :exec
UPDATE product_stock
SET deleted_at = NOW()
WHERE product_id = $1 AND stock_id = $2;

-- name: ListProductStocks :many
SELECT * 
FROM product_stock 
WHERE product_id = $1 AND deleted_at IS NULL 
ORDER BY created_at
LIMIT $2 OFFSET $3; 

-- name: ListProductsInStock :many
SELECT p.*
FROM products p
JOIN product_stock ps ON p.id = ps.product_id
WHERE ps.stock_id = $1 AND ps.deleted_at IS NULL AND p.deleted_at IS NULL
ORDER BY p.name
LIMIT $2 OFFSET $3;