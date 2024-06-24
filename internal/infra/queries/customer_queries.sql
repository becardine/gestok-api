-- name: GetCustomer :one
SELECT * FROM customers WHERE id = ? AND deleted_at IS NULL;

-- name: CreateCustomer :exec
INSERT INTO customers (id, name, email, password, address, phone, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateCustomer :exec
UPDATE customers
SET name = ?, email = ?, address = ?, phone = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: DeleteCustomer :exec
UPDATE customers
SET deleted_at = NOW()
WHERE id = ?;

-- name: ListCustomers :many
SELECT *
FROM customers
WHERE deleted_at IS NULL
ORDER BY name
LIMIT ? OFFSET ?;

-- name: GetCustomerOrders :many
SELECT o.*
FROM customers c
JOIN orders o ON c.id = o.customer_id
WHERE c.id = ? AND c.deleted_at IS NULL AND o.deleted_at IS NULL
ORDER BY o.order_at DESC
LIMIT ? OFFSET ?;

-- name: GetCustomerPayments :many
SELECT p.*
FROM customers c
JOIN payments p ON c.id = p.customer_id
WHERE c.id = ? AND c.deleted_at IS NULL AND p.deleted_at IS NULL
ORDER BY p.payment_at DESC
LIMIT ? OFFSET ?;

-- name: GetCustomerDeliveries :many
SELECT d.*
FROM customers c
JOIN deliveries d ON c.id = d.customer_id
WHERE c.id = ? AND c.deleted_at IS NULL AND d.deleted_at IS NULL
ORDER BY d.delivery_at DESC
LIMIT ? OFFSET ?;

-- name: GetCustomerFeedbacks :many
SELECT f.*
FROM customers c
JOIN feedbacks f ON c.id = f.customer_id
WHERE c.id = ? AND c.deleted_at IS NULL AND f.deleted_at IS NULL
ORDER BY f.created_at DESC
LIMIT ? OFFSET ?;
