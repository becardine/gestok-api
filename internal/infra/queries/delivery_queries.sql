-- name: GetDelivery :one
SELECT * FROM deliveries WHERE id = $1 AND deleted_at IS NULL;

-- name: CreateDelivery :exec
INSERT INTO deliveries (id, order_id, customer_id, delivery_type, delivery_date, delivery_status, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: UpdateDelivery :exec
UPDATE deliveries
SET delivery_type = $2, delivery_date = $3, delivery_status = $4, updated_date = $5
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteDelivery :exec
UPDATE deliveries
SET deleted_at = NOW()
WHERE id = $1;

-- name: ListDeliveries :many
SELECT *
FROM deliveries
WHERE deleted_at IS NULL
ORDER BY delivery_date DESC
    LIMIT $1 OFFSET $2;

-- name: GetDeliveriesByOrderId :many
SELECT *
FROM deliveries
WHERE order_id = $1 AND deleted_at IS NULL
ORDER BY delivery_date DESC
    LIMIT $2 OFFSET $3;

-- name: GetDeliveriesByCustomerId :many
SELECT *
FROM deliveries
WHERE customer_id = $1 AND deleted_at IS NULL
ORDER BY delivery_date DESC
    LIMIT $2 OFFSET $3;