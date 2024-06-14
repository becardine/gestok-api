-- name: GetOrder :one
SELECT * FROM orders WHERE id = $1 AND deleted_at IS NULL;

-- name: CreateOrder :exec
INSERT INTO orders (id, customer_id, order_date, order_status, total_value, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateOrder :exec
UPDATE orders
SET order_date = $2, order_status = $3, total_value = $4, updated_date = $5
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteOrder :exec
UPDATE orders
SET deleted_at = NOW()
WHERE id = $1;

-- name: ListOrders :many
SELECT *
FROM orders
WHERE deleted_at IS NULL
ORDER BY order_date DESC
    LIMIT $1 OFFSET $2;

-- name: GetOrderProducts :many
SELECT p.*
FROM orders o
         JOIN order_products op ON o.id = op.order_id
         JOIN products p ON op.product_id = p.id
WHERE o.id = $1 AND o.deleted_at IS NULL AND p.deleted_at IS NULL
ORDER BY p.name;

-- name: AddOrderProduct :exec
INSERT INTO order_products (order_id, product_id, quantity, unit_price)
VALUES ($1, $2, $3, $4)
    ON CONFLICT (order_id, product_id) DO NOTHING;

-- name: RemoveOrderProduct :execresult
DELETE FROM order_products
WHERE order_id = $1 AND product_id = $2
    RETURNING *;

-- name: UpdateOrderStatus :exec
UPDATE orders
SET order_status = $2
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetOrderByCustomerId :many
SELECT *
FROM orders
WHERE customer_id = $1 AND deleted_at IS NULL
ORDER BY order_date DESC
    LIMIT $2 OFFSET $3;

-- name: GetOrderByProductId :many
SELECT o.*
FROM orders o
         JOIN order_products op ON o.id = op.order_id
WHERE op.product_id = $1 AND o.deleted_at IS NULL
ORDER BY o.order_date DESC
    LIMIT $2 OFFSET $3;