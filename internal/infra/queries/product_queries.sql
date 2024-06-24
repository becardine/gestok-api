-- name: GetProduct :one
SELECT * FROM products WHERE id = ? AND deleted_at IS NULL;

-- name: CreateProduct :exec
INSERT INTO products (id, name, description, price, quantity_in_stock, image_url, category_id, brand_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateProduct :exec
UPDATE products
SET name = ?, description = ?, price = ?, quantity_in_stock = ?, image_url = ?, category_id = ?, brand_id = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: DeleteProduct :exec
UPDATE products
SET deleted_at = NOW()
WHERE id = ?;

-- name: ListProducts :many
SELECT *
FROM products
WHERE deleted_at IS NULL
ORDER BY name
LIMIT ? OFFSET ?;  

-- name: GetProductStocks :many
SELECT s.*
FROM products p
JOIN product_stocks ps ON p.id = ps.product_id
JOIN stocks s ON ps.stock_id = s.id
WHERE p.id = ? AND p.deleted_at IS NULL AND s.deleted_at IS NULL
ORDER BY s.name
LIMIT ? OFFSET ?;

-- name: AddProductStock :exec
INSERT INTO product_stocks (stock_id, product_id)
VALUES (?, ?)
ON DUPLICATE KEY UPDATE stock_id = stock_id;

-- name: RemoveProductStock :execresult
DELETE FROM product_stocks
WHERE stock_id = ? AND product_id = ?;
