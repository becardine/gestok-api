-- name: GetProductStocks :many
SELECT * FROM product_stocks WHERE stock_id = $1;

-- name: GetProductStocksByProductId :many
SELECT * FROM product_stocks WHERE product_id = $1;

-- name: CreateProductStock :exec
INSERT INTO product_stocks (stock_id, product_id)
VALUES ($1, $2);

-- name: DeleteProductStock :exec
DELETE FROM product_stocks WHERE stock_id = $1 AND product_id = $2;

-- name: GetProductStock :one
SELECT * FROM product_stocks WHERE stock_id = $1 AND product_id = $2;