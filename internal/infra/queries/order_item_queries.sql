-- name: CreateOrderItem :one
INSERT INTO order_items (
    id, order_id, product_id, quantity, unit_price, created_at, updated_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetOrderItem :one
SELECT * FROM order_items WHERE id = $1 AND deleted_at IS NULL;

-- name: UpdateOrderItem :exec
UPDATE order_items
SET quantity = $2, unit_price = $3, updated_at = $4
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteOrderItem :exec
UPDATE order_items
SET deleted_at = NOW()
WHERE id = $1;

-- name: ListOrderItems :many
SELECT * 
FROM order_items 
WHERE order_id = $1 AND deleted_at IS NULL 
ORDER BY created_at
LIMIT $2 OFFSET $3; 