-- name: GetCustomer :one
SELECT * FROM customers WHERE id = $1 AND deleted_at IS NULL;

-- name: CreateCustomer :exec
INSERT INTO customers (id, name, email, password, address, phone, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: UpdateCustomer :exec
UPDATE customers
SET name = $2, email = $3, password = $4, address = $5, phone = $6, updated_date = $7
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteCustomer :exec
UPDATE customers
SET deleted_at = NOW()
WHERE id = $1;

-- name: ListCustomers :many
SELECT * FROM customers WHERE deleted_at IS NULL ORDER BY name;

-- name: GetCustomerOrders :many
SELECT o.*
FROM customers c
JOIN orders o ON c.id = o.customer_id
WHERE c.id = $1 AND c.deleted_at IS NULL
ORDER BY o.order_date DESC;

-- name: GetCustomerPayments :many
SELECT p.*
FROM customers c
JOIN payments p ON c.id = p.customer_id
WHERE c.id = $1 AND c.deleted_at IS NULL
ORDER BY p.payment_date DESC;

-- name: GetCustomerDeliveries :many
SELECT d.*
FROM customers c
JOIN deliveries d ON c.id = d.customer_id
WHERE c.id = $1 AND c.deleted_at IS NULL
ORDER BY d.delivery_date DESC;

-- name: GetCustomerFeedbacks :many
SELECT f.*
FROM customers c
JOIN feedbacks f ON c.id = f.customer_id
WHERE c.id = $1 AND c.deleted_at IS NULL
ORDER BY f.created_date DESC;