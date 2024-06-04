-- name: GetDelivery :one
SELECT * FROM deliveries WHERE id = $1 AND deleted_at IS NULL;

-- name: CreateDelivery :exec
INSERT INTO deliveries (id, order_id, customer_id, delivery_type, delivery_date, delivery_status, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: UpdateDelivery :exec
UPDATE deliveries
SET order_id = $2, customer_id = $3, delivery_type = $4, delivery_date = $5, delivery_status = $6, updated_date = $7
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteDelivery :exec
UPDATE deliveries
SET deleted_at = NOW()
WHERE id = $1;

-- name: ListDeliveries :many
SELECT * FROM deliveries WHERE deleted_at IS NULL ORDER BY delivery_date DESC;

-- name: GetDeliveriesByOrderId :many
SELECT *
FROM deliveries
WHERE order_id = $1 AND deleted_at IS NULL
ORDER BY delivery_date DESC;

-- name: GetDeliveriesByCustomerId :many
SELECT *
FROM deliveries
WHERE customer_id = $1 AND deleted_at IS NULL
ORDER BY delivery_date DESC;