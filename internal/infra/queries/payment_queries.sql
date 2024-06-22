-- name: GetPayment :one
SELECT * FROM payments WHERE id = ? AND deleted_at IS NULL;

-- name: CreatePayment :exec
INSERT INTO payments (id, order_id, customer_id, payment_type, payment_at, payment_value, payment_status, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdatePayment :exec
UPDATE payments
SET payment_type = ?, payment_at = ?, payment_value = ?, payment_status = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: DeletePayment :exec
UPDATE payments
SET deleted_at = NOW()
WHERE id = ?;

-- name: ListPayments :many
SELECT *
FROM payments
WHERE deleted_at IS NULL
ORDER BY payment_at DESC
LIMIT 1 OFFSET 2;

-- name: GetPaymentsByOrderId :many
SELECT *
FROM payments
WHERE order_id = ? AND deleted_at IS NULL
ORDER BY payment_at DESC
LIMIT 2 OFFSET 3;

-- name: GetPaymentsByCustomerId :many
SELECT *
FROM payments
WHERE customer_id = ? AND deleted_at IS NULL
ORDER BY payment_at DESC
LIMIT 2 OFFSET 3;
