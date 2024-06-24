// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: payment_queries.sql

package sqlc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createPayment = `-- name: CreatePayment :exec
INSERT INTO payments (id, order_id, customer_id, payment_type, payment_at, payment_value, payment_status, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreatePaymentParams struct {
	ID            uuid.UUID
	OrderID       uuid.UUID
	CustomerID    uuid.UUID
	PaymentType   string
	PaymentAt     sql.NullTime
	PaymentValue  float64
	PaymentStatus string
	CreatedAt     sql.NullTime
	UpdatedAt     sql.NullTime
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) error {
	_, err := q.db.ExecContext(ctx, createPayment,
		arg.ID,
		arg.OrderID,
		arg.CustomerID,
		arg.PaymentType,
		arg.PaymentAt,
		arg.PaymentValue,
		arg.PaymentStatus,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deletePayment = `-- name: DeletePayment :exec
UPDATE payments
SET deleted_at = NOW()
WHERE id = ?
`

func (q *Queries) DeletePayment(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePayment, id)
	return err
}

const getPayment = `-- name: GetPayment :one
SELECT id, order_id, customer_id, payment_type, payment_at, payment_value, payment_status, deleted_at, created_at, updated_at FROM payments WHERE id = ? AND deleted_at IS NULL
`

func (q *Queries) GetPayment(ctx context.Context, id uuid.UUID) (Payment, error) {
	row := q.db.QueryRowContext(ctx, getPayment, id)
	var i Payment
	err := row.Scan(
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
	)
	return i, err
}

const getPaymentsByCustomerId = `-- name: GetPaymentsByCustomerId :many
SELECT id, order_id, customer_id, payment_type, payment_at, payment_value, payment_status, deleted_at, created_at, updated_at
FROM payments
WHERE customer_id = ? AND deleted_at IS NULL
ORDER BY payment_at DESC
LIMIT 2 OFFSET 3
`

func (q *Queries) GetPaymentsByCustomerId(ctx context.Context, customerID uuid.UUID) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, getPaymentsByCustomerId, customerID)
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

const getPaymentsByOrderId = `-- name: GetPaymentsByOrderId :many
SELECT id, order_id, customer_id, payment_type, payment_at, payment_value, payment_status, deleted_at, created_at, updated_at
FROM payments
WHERE order_id = ? AND deleted_at IS NULL
ORDER BY payment_at DESC
LIMIT 2 OFFSET 3
`

func (q *Queries) GetPaymentsByOrderId(ctx context.Context, orderID uuid.UUID) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, getPaymentsByOrderId, orderID)
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

const listPayments = `-- name: ListPayments :many
SELECT id, order_id, customer_id, payment_type, payment_at, payment_value, payment_status, deleted_at, created_at, updated_at
FROM payments
WHERE deleted_at IS NULL
ORDER BY payment_at DESC
LIMIT 1 OFFSET 2
`

func (q *Queries) ListPayments(ctx context.Context) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, listPayments)
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

const updatePayment = `-- name: UpdatePayment :exec
UPDATE payments
SET payment_type = ?, payment_at = ?, payment_value = ?, payment_status = ?, updated_at = ?
WHERE id = ? AND deleted_at IS NULL
`

type UpdatePaymentParams struct {
	PaymentType   string
	PaymentAt     sql.NullTime
	PaymentValue  float64
	PaymentStatus string
	UpdatedAt     sql.NullTime
	ID            uuid.UUID
}

func (q *Queries) UpdatePayment(ctx context.Context, arg UpdatePaymentParams) error {
	_, err := q.db.ExecContext(ctx, updatePayment,
		arg.PaymentType,
		arg.PaymentAt,
		arg.PaymentValue,
		arg.PaymentStatus,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
