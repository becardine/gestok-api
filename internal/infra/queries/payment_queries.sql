-- name: GetPayment :one
SELECT * FROM payments WHERE id = $1 AND deleted_at IS NULL;

-- name: CreatePayment :exec
INSERT INTO payments (id, order_id, customer_id, payment_type, payment_date, payment_value, payment_status, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: UpdatePayment :exec
UPDATE payments
SET payment_type = $2, payment_date = $3, payment_value = $4, payment_status = $5, updated_date = $6
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeletePayment :exec
UPDATE payments
SET deleted_at = NOW()
WHERE id = $1;

-- name: ListPayments :many
SELECT *
FROM payments
WHERE deleted_at IS NULL
ORDER BY payment_date DESC
    LIMIT $1 OFFSET $2;

-- name: GetPaymentsByOrderId :many
SELECT *
FROM payments
WHERE order_id = $1 AND deleted_at IS NULL
ORDER BY payment_date DESC
    LIMIT $2 OFFSET $3;

-- name: GetPaymentsByCustomerId :many
SELECT *
FROM payments
WHERE customer_id = $1 AND deleted_at IS NULL
ORDER BY payment_date DESC
    LIMIT $2 OFFSET $3;