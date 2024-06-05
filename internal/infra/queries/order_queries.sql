-- name: GetOrder :one
SELECT * FROM orders WHERE id = $1 AND deleted_at IS NULL;

-- name: CreateOrder :exec
INSERT INTO orders (id, customer_id, order_date, order_status, total_value, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateOrder :exec
UPDATE orders
SET customer_id = $2, order_date = $3, order_status = $4, total_value = $5, updated_date = $6
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteOrder :exec
UPDATE orders
SET deleted_at = NOW()
WHERE id = $1;

-- name: ListOrders :many
SELECT * FROM orders WHERE deleted_at IS NULL ORDER BY order_date DESC;

-- name: GetOrderProducts :many
SELECT p.* 
FROM orders o
JOIN order_products op ON o.id = op.order_id
JOIN products p ON op.product_id = p.id
WHERE o.id = $1 AND o.deleted_at IS NULL;

-- name: AddOrderProduct :exec
INSERT INTO order_products (order_id, product_id, quantity, unit_price)
VALUES ($1, $2, $3, $4);

-- name: RemoveOrderProduct :exec
DELETE FROM order_products WHERE order_id = $1 AND product_id = $2;

-- name: UpdateOrderStatus :exec
UPDATE orders
SET order_status = $2
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetOrderByCustomerId :many
SELECT *
FROM orders
WHERE customer_id = $1 AND deleted_at IS NULL
ORDER BY order_date DESC;

-- name: GetOrderByProductId :many
SELECT o.*
FROM orders o
JOIN order_products op ON o.id = op.order_id
WHERE op.product_id = $1 AND o.deleted_at IS NULL
ORDER BY o.order_date DESC;