// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: customer_queries.sql

package sqlc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createCustomer = `-- name: CreateCustomer :exec
INSERT INTO customers (id, name, email, password, address, phone, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateCustomerParams struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	Address   sql.NullString
	Phone     sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) error {
	_, err := q.db.ExecContext(ctx, createCustomer,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Address,
		arg.Phone,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteCustomer = `-- name: DeleteCustomer :exec
UPDATE customers
SET deleted_at = NOW()
WHERE id = ?
`

func (q *Queries) DeleteCustomer(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCustomer, id)
	return err
}

const getCustomer = `-- name: GetCustomer :one
SELECT id, name, email, password, address, phone, deleted_at, created_at, updated_at FROM customers WHERE id = ? AND deleted_at IS NULL
`

func (q *Queries) GetCustomer(ctx context.Context, id uuid.UUID) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomer, id)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Address,
		&i.Phone,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCustomerDeliveries = `-- name: GetCustomerDeliveries :many
SELECT d.id, d.order_id, d.customer_id, d.delivery_type, d.delivery_at, d.delivery_status, d.deleted_at, d.created_at, d.updated_at
FROM customers c
JOIN deliveries d ON c.id = d.customer_id
WHERE c.id = ? AND c.deleted_at IS NULL AND d.deleted_at IS NULL
ORDER BY d.delivery_at DESC
LIMIT ? OFFSET ?
`

type GetCustomerDeliveriesParams struct {
	ID     uuid.UUID
	Limit  int32
	Offset int32
}

func (q *Queries) GetCustomerDeliveries(ctx context.Context, arg GetCustomerDeliveriesParams) ([]Delivery, error) {
	rows, err := q.db.QueryContext(ctx, getCustomerDeliveries, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Delivery
	for rows.Next() {
		var i Delivery
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.CustomerID,
			&i.DeliveryType,
			&i.DeliveryAt,
			&i.DeliveryStatus,
			&i.DeletedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCustomerFeedbacks = `-- name: GetCustomerFeedbacks :many
SELECT f.id, f.customer_id, f.order_id, f.rating, f.comment, f.deleted_at, f.created_at, f.updated_at
FROM customers c
JOIN feedbacks f ON c.id = f.customer_id
WHERE c.id = ? AND c.deleted_at IS NULL AND f.deleted_at IS NULL
ORDER BY f.created_at DESC
LIMIT ? OFFSET ?
`

type GetCustomerFeedbacksParams struct {
	ID     uuid.UUID
	Limit  int32
	Offset int32
}

func (q *Queries) GetCustomerFeedbacks(ctx context.Context, arg GetCustomerFeedbacksParams) ([]Feedback, error) {
	rows, err := q.db.QueryContext(ctx, getCustomerFeedbacks, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feedback
	for rows.Next() {
		var i Feedback
		if err := rows.Scan(
			&i.ID,
			&i.CustomerID,
			&i.OrderID,
			&i.Rating,
			&i.Comment,
			&i.DeletedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCustomerOrders = `-- name: GetCustomerOrders :many
SELECT o.id, o.customer_id, o.order_at, o.order_status, o.total_value, o.deleted_at, o.created_at, o.updated_at
FROM customers c
JOIN orders o ON c.id = o.customer_id
WHERE c.id = ? AND c.deleted_at IS NULL AND o.deleted_at IS NULL
ORDER BY o.order_at DESC
LIMIT ? OFFSET ?
`

type GetCustomerOrdersParams struct {
	ID     uuid.UUID
	Limit  int32
	Offset int32
}

func (q *Queries) GetCustomerOrders(ctx context.Context, arg GetCustomerOrdersParams) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getCustomerOrders, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.CustomerID,
			&i.OrderAt,
			&i.OrderStatus,
			&i.TotalValue,
			&i.DeletedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCustomerPayments = `-- name: GetCustomerPayments :many
SELECT p.id, p.order_id, p.customer_id, p.payment_type, p.payment_at, p.payment_value, p.payment_status, p.deleted_at, p.created_at, p.updated_at
FROM customers c
JOIN payments p ON c.id = p.customer_id
WHERE c.id = ? AND c.deleted_at IS NULL AND p.deleted_at IS NULL
ORDER BY p.payment_at DESC
LIMIT ? OFFSET ?
`

type GetCustomerPaymentsParams struct {
	ID     uuid.UUID
	Limit  int32
	Offset int32
}

func (q *Queries) GetCustomerPayments(ctx context.Context, arg GetCustomerPaymentsParams) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, getCustomerPayments, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Payment
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.CustomerID,
			&i.PaymentType,
			&i.PaymentAt,
			&i.PaymentValue,
			&i.PaymentStatus,
			&i.DeletedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCustomers = `-- name: ListCustomers :many
SELECT id, name, email, password, address, phone, deleted_at, created_at, updated_at
FROM customers
WHERE deleted_at IS NULL
ORDER BY name
LIMIT ? OFFSET ?
`

type ListCustomersParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListCustomers(ctx context.Context, arg ListCustomersParams) ([]Customer, error) {
	rows, err := q.db.QueryContext(ctx, listCustomers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Customer
	for rows.Next() {
		var i Customer
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.Address,
			&i.Phone,
			&i.DeletedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCustomer = `-- name: UpdateCustomer :exec
UPDATE customers
SET name = ?, email = ?, address = ?, phone = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL
`

type UpdateCustomerParams struct {
	Name      string
	Email     string
	Address   sql.NullString
	Phone     sql.NullString
	UpdatedAt sql.NullTime
	ID        uuid.UUID
}

func (q *Queries) UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) error {
	_, err := q.db.ExecContext(ctx, updateCustomer,
		arg.Name,
		arg.Email,
		arg.Address,
		arg.Phone,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
