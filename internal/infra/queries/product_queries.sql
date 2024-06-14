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
SELECT *
FROM products
WHERE deleted_at IS NULL
ORDER BY name
LIMIT $1 OFFSET $2;

-- name: GetProductStocks :many
SELECT s.*
FROM products p
         JOIN product_stocks ps ON p.id = ps.product_id
         JOIN stocks s ON ps.stock_id = s.id
WHERE p.id = $1 AND p.deleted_at IS NULL AND s.deleted_at IS NULL
ORDER BY s.name;

-- name: AddProductStock :exec
INSERT INTO product_stocks (stock_id, product_id)
VALUES ($1, $2)
ON CONFLICT (stock_id, product_id) DO NOTHING;

-- name: RemoveProductStock :execresult
DELETE FROM product_stocks ps
WHERE ps.stock_id = $1 AND ps.product_id = $2
RETURNING CASE WHEN EXISTS(SELECT 1 FROM product_stocks WHERE stock_id = $1 AND product_id = $2) THEN 0 ELSE 1 END;