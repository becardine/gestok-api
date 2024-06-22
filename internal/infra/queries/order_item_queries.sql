-- name: CreateOrderItem :exec
INSERT INTO order_items (
    id, order_id, product_id, quantity, unit_price, created_at, updated_at
)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetOrderItem :one
SELECT * FROM order_items WHERE id = ? AND deleted_at IS NULL;

-- name: UpdateOrderItem :exec
UPDATE order_items
SET quantity = ?, unit_price = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: DeleteOrderItem :exec
UPDATE order_items
SET deleted_at = NOW()
WHERE id = ?;

-- name: ListOrderItems :many
SELECT * 
FROM order_items 
WHERE order_id = ? AND deleted_at IS NULL 
ORDER BY created_at
LIMIT 2 OFFSET 3;
