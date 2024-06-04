-- name: GetProduct :one
SELECT * FROM products WHERE id = $1;

-- name: CreateProduct :exec
INSERT INTO products (name, description, price, quantity_in_stock, image_url, category_id, brand_id, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: UpdateProduct :exec
UPDATE products
SET name = $2, description = $3, price = $4, quantity_in_stock = $5, image_url = $6, category_id = $7, brand_id = $8, updated_date = $9
WHERE id = $1;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;

-- name: ListProducts :many
SELECT * FROM products ORDER BY name;