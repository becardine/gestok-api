-- name: GetProduct :one
SELECT * FROM products WHERE id = $1 AND deleted_at IS NULL;

-- name: CreateProduct :exec
INSERT INTO products (id, name, description, price, quantity_in_stock, image_url, category_id, brand_id, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

-- name: UpdateProduct :exec
UPDATE products
SET name = $2, description = $3, price = $4, quantity_in_stock = $5, image_url = $6, category_id = $7, brand_id = $8, updated_date = $9
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteProduct :exec
UPDATE products
SET deleted_at = NOW()
WHERE id = $1;

-- name: ListProducts :many
SELECT * FROM products WHERE deleted_at IS NULL ORDER BY name;

-- name: GetProductProducts :many
SELECT p.*
FROM products s
JOIN product_stocks ps ON s.id = ps.product_id
JOIN stocks p ON ps.stock_id = p.id
WHERE s.id = $1 AND s.deleted_at IS NULL;

-- name: AddProductStock :exec
INSERT INTO product_stocks (stock_id, product_id)
VALUES ($1, $2);

-- name: RemoveProductStock :exec
DELETE FROM product_stocks WHERE stock_id = $1 AND product_id = $2;