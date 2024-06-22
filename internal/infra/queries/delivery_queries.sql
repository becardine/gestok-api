-- name: GetDelivery :one
SELECT * FROM deliveries WHERE id = ? AND deleted_at IS NULL;

-- name: CreateDelivery :exec
INSERT INTO deliveries (id, order_id, customer_id, delivery_type, delivery_at, delivery_status, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateDelivery :exec
UPDATE deliveries
SET delivery_type = ?, delivery_at = ?, delivery_status = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: DeleteDelivery :exec
UPDATE deliveries
SET deleted_at = NOW()
WHERE id = ?;

-- name: ListDeliveries :many
SELECT *
FROM deliveries
WHERE deleted_at IS NULL
ORDER BY delivery_at DESC
LIMIT 1 OFFSET 2;

-- name: GetDeliveriesByOrderId :many
SELECT *
FROM deliveries
WHERE order_id = ? AND deleted_at IS NULL
ORDER BY delivery_at DESC
LIMIT 2 OFFSET 3;

-- name: GetDeliveriesByCustomerId :many
SELECT *
FROM deliveries
WHERE customer_id = ? AND deleted_at IS NULL
ORDER BY delivery_at DESC
LIMIT 2 OFFSET 3;
