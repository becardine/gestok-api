-- name: GetCustomer :one
SELECT * FROM customers WHERE id = $1 AND deleted_at IS NULL;

-- name: CreateCustomer :exec
INSERT INTO customers (id, name, email, password, address, phone, created_date, updated_date)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: UpdateCustomer :exec
UPDATE customers
SET name = $2, email = $3, address = $4, phone = $5, updated_date = $6
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteCustomer :exec
UPDATE customers
SET deleted_at = NOW()
WHERE id = $1;

-- name: ListCustomers :many
SELECT *
FROM customers
WHERE deleted_at IS NULL
ORDER BY name
    LIMIT $1 OFFSET $2;

-- name: GetCustomerOrders :many
SELECT o.*
FROM customers c
         JOIN orders o ON c.id = o.customer_id
WHERE c.id = $1 AND c.deleted_at IS NULL AND o.deleted_at IS NULL
ORDER BY o.order_date DESC
    LIMIT $2 OFFSET $3;

-- name: GetCustomerPayments :many
SELECT p.*
FROM customers c
         JOIN payments p ON c.id = p.customer_id
WHERE c.id = $1 AND c.deleted_at IS NULL AND p.deleted_at IS NULL
ORDER BY p.payment_date DESC
    LIMIT $2 OFFSET $3;

-- name: GetCustomerDeliveries :many
SELECT d.*
FROM customers c
         JOIN deliveries d ON c.id = d.customer_id
WHERE c.id = $1 AND c.deleted_at IS NULL AND d.deleted_at IS NULL
ORDER BY d.delivery_date DESC
    LIMIT $2 OFFSET $3;

-- name: GetCustomerFeedbacks :many
SELECT f.*
FROM customers c
         JOIN feedbacks f ON c.id = f.customer_id
WHERE c.id = $1 AND c.deleted_at IS NULL AND f.deleted_at IS NULL
ORDER BY f.created_date DESC
    LIMIT $2 OFFSET $3;