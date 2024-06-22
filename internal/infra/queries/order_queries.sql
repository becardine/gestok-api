-- name: GetOrder :one
SELECT * FROM orders WHERE id = ? AND deleted_at IS NULL;

-- name: CreateOrder :exec
INSERT INTO orders (id, customer_id, order_at, order_status, total_value, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateOrder :exec
UPDATE orders
SET order_at = ?, order_status = ?, total_value = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: DeleteOrder :exec
UPDATE orders
SET deleted_at = NOW()
WHERE id = ?;

-- name: ListOrders :many
SELECT *
FROM orders
WHERE deleted_at IS NULL
ORDER BY order_at DESC
LIMIT 1 OFFSET 2;

-- name: GetOrderProducts :many
SELECT p.*
FROM orders o
JOIN order_products op ON o.id = op.order_id
JOIN products p ON op.product_id = p.id
WHERE o.id = ? AND o.deleted_at IS NULL AND p.deleted_at IS NULL
ORDER BY p.name;

-- name: AddOrderProduct :exec
INSERT INTO order_products (order_id, product_id, quantity, unit_price)
VALUES (?, ?, ?, ?)
ON DUPLICATE KEY UPDATE order_id = order_id;

-- name: RemoveOrderProduct :execresult
DELETE FROM order_products
WHERE order_id = ? AND product_id = ?;

-- name: UpdateOrderStatus :exec
UPDATE orders
SET order_status = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: GetOrderByCustomerId :many
SELECT *
FROM orders
WHERE customer_id = ? AND deleted_at IS NULL
ORDER BY order_at DESC
LIMIT 2 OFFSET 3;

-- name: GetOrderByProductId :many
SELECT o.*
FROM orders o
JOIN order_products op ON o.id = op.order_id
WHERE op.product_id = ? AND o.deleted_at IS NULL
ORDER BY o.order_at DESC
LIMIT 2 OFFSET 3;
